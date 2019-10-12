package models

import (
	"context"
	"golang.org/x/net/webdav"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type WebdavFs struct {
	rootPath       string//根目录
	prefix 			string //请求前缀
	requestRoot     string //请求根目录
	User        *User     //用户
	Context     context.Context //请求
	webdavFile *WebdavFileSystem
	ShareRootList []Share  //根目录下的分享文件列表
	ShareInfo *Share //请求根目录的分享信息
}

func (this WebdavFs) resolve(name string) string  {
	if filepath.Separator != '/' && strings.IndexRune(name, filepath.Separator) >= 0 ||
		strings.Contains(name, "\x00") {
		return ""
	}
	path := this.rootPath
	if path == "" {
		path = "."
	}
	return filepath.Join(path,filepath.FromSlash(slashClean(name)))
}
func (this WebdavFs) Mkdir(ctx context.Context, name string, perm os.FileMode) error{
	if name = this.resolve(name); name == "" {
		return os.ErrNotExist
	}
	return os.Mkdir(name, perm)
}
func (this WebdavFs) OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) (webdav.File, error){
	webdavFile := WebdavFileSystem{
		prefix:this.prefix,
		rootPath:this.rootPath,
		userRoot:this.User.Username,
		path:name,
		requestRoot:this.requestRoot,
		ShareRootList:this.ShareRootList,
		ShareInfo:this.ShareInfo,
	}
	return webdavFile, nil
}
func (this WebdavFs) RemoveAll(ctx context.Context, name string) error{
	if name = this.resolve(name); name == "" {
		return os.ErrNotExist
	}
	if name == filepath.Clean(this.rootPath) {
		// Prohibit removing the virtual root directory.
		return os.ErrInvalid
	}
	return os.RemoveAll(name)
}
func (this WebdavFs) Rename(ctx context.Context, oldName, newName string) error{
	if oldName = this.resolve(oldName); oldName == "" {
		return os.ErrNotExist
	}
	if newName = this.resolve(newName); newName == "" {
		return os.ErrNotExist
	}
	if root := filepath.Clean(this.rootPath); root == oldName || root == newName {
		// Prohibit renaming from or to the virtual root directory.
		return os.ErrInvalid
	}
	return os.Rename(oldName, newName)
}
func (this *WebdavFs) Stat(ctx context.Context, name string) (os.FileInfo, error){
	webdavFile := WebdavFileSystem{
		prefix:this.prefix,
		rootPath:this.rootPath,
		userRoot:this.User.Username,
		path:name,
		requestRoot:this.requestRoot,
		ShareRootList:this.ShareRootList,
		ShareInfo:this.ShareInfo,
	}
	return webdavFile.Stat()
}
/**
rootPath 根目录, prefix 请求前缀,requestRoot 请求根目录， User 用户, Context context.Context 请求信息
 */
func NewWebdavFs( rootPath string, prefix string,requestRoot string, User User, Context context.Context) *WebdavFs {
	//获取根目录下的分享文件列表
	ShareRootList,_ := ShareRootList(requestRoot)
	//分享文件及其子目录
	shareInfo, err :=ShareInfo(requestRoot)

	webdavFs := WebdavFs{
		prefix: prefix,
		rootPath:rootPath,
		requestRoot:requestRoot,
		User: &User,
		Context:Context,
		ShareRootList: ShareRootList,
	}
	if err == nil {
		webdavFs.ShareInfo = &shareInfo
	}
	return &webdavFs
}
/*
* Implement a webdav.File and os.Stat : https://godoc.org/golang.org/x/net/webdav#File
*/
type WebdavFileSystem struct {
	prefix string //请求前缀
	rootPath string //根目录
	userRoot string //用户根目录
	path    string //请求目录
	requestRoot     string //请求根目录
	fread   *os.File
	fwrite  *os.File

	ShareRootList []Share  //根目录下的分享文件列表
	ShareInfo *Share //请求根目录的分享信息
}
func (this WebdavFileSystem) resolve(name string) string  {
	if filepath.Separator != '/' && strings.IndexRune(name, filepath.Separator) >= 0 ||
		strings.Contains(name, "\x00") {
		return ""
	}
	rootPath := this.rootPath
	if rootPath == "" {
		rootPath = "."
	}
	for _,shareFile := range this.ShareRootList {//根目录的分享文件
		if name == shareFile.ItemTarget {
			return filepath.Join(rootPath,shareFile.FileSource,filepath.FromSlash(slashClean(shareFile.FileTarget)))
		}
	}
	if this.ShareInfo != nil{//是分享目录及子目录
		name = strings.Replace(name, this.ShareInfo.ItemTarget, this.ShareInfo.FileTarget, 1)
		return filepath.Join(rootPath,this.ShareInfo.FileSource,name)

	}
	return filepath.Join(rootPath,this.userRoot,filepath.FromSlash(slashClean(name)))
}
func (this WebdavFileSystem) Read(p []byte) (n int, err error) {
	if strings.HasPrefix(filepath.Base(this.path), ".") {
		return 0, os.ErrNotExist
	}
	if this.fread == nil {
		return -1, os.ErrInvalid
	}
	return this.fread.Read(p)
}
func (this WebdavFileSystem) Close() error {
	if this.fread != nil {
		if this.fread.Close() == nil {
			this.fread = nil
		}
	}
	if this.fwrite != nil {

		if this.fwrite.Close() == nil {
			this.fwrite = nil
		}
	}
	return nil
}

func (this WebdavFileSystem) Seek(offset int64, whence int) (int64, error) {
	if this.fread == nil {
		return offset, os.ErrNotExist
	}
	a, err := this.fread.Seek(offset, whence)
	if err != nil {
		return a, os.ErrNotExist
	}
	return a, nil
}
func (this WebdavFileSystem) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	name = this.resolve(name)
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f,nil
}
func (this WebdavFileSystem) Readdir(count int) (files []os.FileInfo, err error) {

	if strings.HasPrefix(filepath.Base(this.path), ".") {
		err = os.ErrNotExist
		return
	}
	path := this.resolve(this.path)
	f, err := ioutil.ReadDir(path)

	for _, fileInfo := range f {
		fs := this
		fs.path = fileInfo.Name()

		files = append(files,fs)
	}

	if this.requestRoot == this.path && len(this.ShareRootList) > 0 {
		for _, shareFile := range this.ShareRootList {
			fs := this
			fs.path = shareFile.ItemTarget
			fs.userRoot = shareFile.FileSource
			files = append(files, fs)
		}
	}
	return
}

func (this WebdavFileSystem) Stat() (os.FileInfo, error) {
	if strings.HasSuffix(this.path, "/") {
		_, err := this.Readdir(0)
		if err != nil {
			return nil, os.ErrNotExist
		}
		return this, nil
	}
	path := this.resolve(this.path)
	if  path == "" {
		return nil, os.ErrNotExist
	}
	baseDir := filepath.Base(path)
	files, err := ioutil.ReadDir(strings.TrimSuffix(path, baseDir))

	if err != nil {
		return nil, os.ErrNotExist
	}
	found := false
	for i := range files {
		if files[i].Name() == baseDir {
			found = true
			break
		}
	}
	if found == false {
		return nil, os.ErrNotExist
	}
	return this, nil
}

func (this WebdavFileSystem) Write(p []byte) (int, error) {
	if this.fwrite == nil {
		return 0, os.ErrNotExist
	}
	if strings.HasPrefix(filepath.Base(this.path), ".") {
		return 0, os.ErrNotExist
	}
	return this.fwrite.Write(p)
}

func (this WebdavFileSystem) Name() string {
	return filepath.Base(this.path)
}

func (this WebdavFileSystem) Size() int64 {
	path := this.resolve(this.path)
	s, err := os.Stat(path)
	if err == nil {
		return s.Size()
	}
	return  0
}

func (this WebdavFileSystem) Mode() os.FileMode {
	path := this.resolve(this.path)
	s, err := os.Stat(path)
	if err == nil {
		return s.Mode()
	}
	return 0
}

func (this WebdavFileSystem) ModTime() time.Time {
	path := this.resolve(this.path)
	s, err := os.Stat(path)
	if err == nil {
		return s.ModTime()
	}
	return time.Now()
}
func (this WebdavFileSystem) IsDir() bool {
	path := this.resolve(this.path)
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func (this WebdavFileSystem) Sys() interface{} {
	path := this.resolve(this.path)
	s, err := os.Stat(path)
	if err == nil {
		return s.Sys()
	}
	return nil
}

/*func (this WebdavFile) ETag(ctx context.Context) (string, error) {
	// Building an etag can be an expensive call if the data isn't available locally.
	// => 2 etags strategies:
	// - use a legit etag value when the data is already in our cache
	// - use a dummy value that's changing all the time when we don't have much info

	etag := Hash(fmt.Sprintf("%d%s", this.ModTime().UnixNano(), this.path), 20)
	if this.fread != nil {
		if s, err := this.fread.Stat(); err == nil {
			etag = Hash(fmt.Sprintf(`"%x%x"`, this.path, s.Size()), 20)
		}
	}
	return etag, nil
}*/