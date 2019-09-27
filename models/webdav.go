package models

import (
	"context"
	"golang.org/x/net/webdav"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	_ "unsafe"
)


type WebdavFile struct {
	path string
	webdav.Dir
	files [] os.FileInfo
}

func (this *WebdavFile) Readdir(count int)([]os.FileInfo,error)  {
	if strings.HasPrefix(filepath.Base(this.path), ".") {
		return nil, os.ErrNotExist
	}

	if this.path != ""{
		this.Dir = webdav.Dir(this.path)
		if files,err := ioutil.ReadDir(this.path); err == nil{
			this.files = files
		}
	}
	return this.files,nil
}

type WebDavFs struct {
	User
	webdav.Dir
	Storage string//文件系统默认是用户名
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
func (d *WebDavFs) SetStartReqName(prefix,p string){
	name := ""
	if prefix == "" {
		name = p
	}
	if r := strings.TrimPrefix(p, prefix); len(r) < len(p) {
		name = r
	}
	name = slashClean(name)
	if shareRecord ,err := ShareInfo(name); err==nil{
		d.Storage = shareRecord.FileSource//分享的用户原始目录
		d.newBaseDir = shareRecord.FileTarget
		d.oldBaseDir = shareRecord.ItemTarget
	}
}

func (d *WebDavFs) resolve(name string) string {
	name = slashClean(name)
	if d.newBaseDir != ""{
		//log.Println("newName1:", name,d.oldBaseDir,d.newBaseDir)
		//log.Println("replace:",strings.Replace(name,d.oldBaseDir,d.newBaseDir,1))
		name =  strings.Replace(name,d.oldBaseDir,d.newBaseDir,1)
		//log.Println("newName:", name,d.oldBaseDir,d.newBaseDir)
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

func (d WebDavFs) OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) ( webdav.File,  error) {
	//log.Println("myWebDav:",name,flag,perm)
	if name = d.resolve(name); name == "" {
		return nil, os.ErrNotExist
	}
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	//st,st_er := f.Stat()
	//log.Println("stateInfo:",st,st_er)
	//di,er := f.Readdir(1)
	//log.Println("dirInfo:",di,er)


	return f, nil
}

func (d *WebDavFs) Stat(ctx context.Context, name string) (os.FileInfo, error) {
	//log.Println("stat:",name)
	if name = d.resolve(name); name == "" {
		return nil, os.ErrNotExist
	}
	//log.Println("stat2:",name)
	return os.Stat(name)
}
