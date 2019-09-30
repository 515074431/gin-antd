package models

import (
	"context"
	//"github.com/515074431/gin-antd/pkg/webdav"
	"golang.org/x/net/webdav"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type WebdavFile struct {
	Dir     string
	Prefix  string //请求的前缀
	path    string //请求路径   绝对路径
	baseDir string //基本请求地址   相对路径
	reqName string //请求名字   相对路径
	*os.File
	files [] os.FileInfo
}

func (this *WebdavFile) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {

	if f, err := os.OpenFile(name, flag, perm); err == nil {
		//this.File = f
		return f, err
	} else {
		return nil, err
	}

}

func (this *WebdavFile) Readdir(count int) ([]os.FileInfo, error) {
	if strings.HasPrefix(filepath.Base(this.path), ".") {
		return nil, os.ErrNotExist
	}

	if this.path != "" {
		//this.File,err  := os.Open(this.path)
		if files, err := ioutil.ReadDir(this.path); err == nil {
			this.files = files
		}
	}
	log.Println("Readdir 请求的目录：", this.path, this.baseDir, this.reqName, this.Name())
	if this.reqName == this.baseDir { //根目录下去获取分享的文件
		log.Println("是根目录请求：", this.baseDir, this.reqName, this.path)
		if shareFiles, err := ShareRootList(this.reqName); err == nil { //找到分享的文件了
			for _, shareFile := range shareFiles {
				log.Println("ShareFile :", shareFile)
				shareFileStr := filepath.Join(this.Dir, shareFile.FileSource, shareFile.FileTarget)
				log.Println("ShareFileStr :", shareFileStr)

				if shareFileInfo, err := os.Stat(shareFileStr); err == nil {
					log.Println("ShareFileInfo :", shareFileInfo)




					this.files = append(this.files, shareFileInfo)
				}

			}
		}

	}
	log.Println("WebdavFile->files", this.files)
	for i, f := range this.files {

		log.Println("WebdavFile->files", i,f)
	}
	return this.files, nil
}

type WebDavFs struct {
	Prefix string
	User
	webdav.Dir
	Storage string //文件系统默认是用户名
	WebdavFile
	oldBaseDir string
	newBaseDir string
}

// slashClean is equivalent to but slightly more efficient than
// path.Clean("/" + name).
func slashClean(name string) string {
	if name == "" || name[0] != '/' {
		name = "/" + name
	}
	return path.Clean(name)
}

//去除前缀 一定要先执行一下
func (d *WebDavFs) SetStartReqName(prefix, p string) {
	name := ""
	if prefix == "" {
		name = p
	}
	if r := strings.TrimPrefix(p, prefix); len(r) < len(p) {
		name = r
	}
	name = slashClean(name)
	if shareRecord, err := ShareInfo(name); err == nil {
		d.Storage = shareRecord.FileSource //分享的用户原始目录
		d.newBaseDir = shareRecord.FileTarget
		d.oldBaseDir = shareRecord.ItemTarget
	} else {
		d.newBaseDir = name
		d.oldBaseDir = name
	}
	d.Prefix = prefix
}

func (d *WebDavFs) resolve(name string) string {
	name = slashClean(name)
	if d.newBaseDir != "" {
		name = strings.Replace(name, d.oldBaseDir, d.newBaseDir, 1)
	}

	// This implementation is based on Dir.Open's code in the standard net/http package.
	if filepath.Separator != '/' && strings.IndexRune(name, filepath.Separator) >= 0 ||
		strings.Contains(name, "\x00") {
		return ""
	}
	dir := string(d.Dir)
	if dir == "" {
		dir = "."
	}
	return filepath.Join(dir, d.Storage, filepath.FromSlash(name))
}

func (d WebDavFs) OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) (webdav.File, error) {
	reqName := slashClean(name) //请求名字 相对路径
	if name = d.resolve(name); name == "" {
		return nil, os.ErrNotExist
	}
	file := &WebdavFile{Prefix: d.Prefix, Dir: string(d.Dir), path: name, baseDir: d.oldBaseDir, reqName: reqName}
	f, err := file.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	file.File = f
	//st,st_er := f.Stat()
	////log.Println("stateInfo:",st,st_er)
	//di,er := f.Readdir(1)
	////log.Println("dirInfo:",di,er)

	return file, nil
}

func (d *WebDavFs) Stat(ctx context.Context, name string) (os.FileInfo, error) {
	if name = d.resolve(name); name == "" {
		return nil, os.ErrNotExist
	}
	return os.Stat(name)
}
