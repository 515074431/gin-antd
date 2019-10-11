package v1

import (
	"errors"
	"github.com/515074431/gin-antd/models"
	"github.com/515074431/gin-antd/pkg/setting"
	"strings"

	//"github.com/515074431/gin-antd/pkg/setting"
	"github.com/515074431/gin-antd/pkg/util"

	//"github.com/515074431/gin-antd/pkg/webdav"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/webdav"
	"log"
)

func WebDav(c *gin.Context) {
	//c.Request.Method = "PROPFIND"
	//fileName := c.Param("file")

	if user, ok := util.GetIdentify(c); ok {
		log.Println("User:", user)
		//fs := &webdav.Server{
		//	webdav.Handler{Prefix: "/api/v1/dav/",
		//		FileSystem: webdav.Dir("d:/tmp/"),
		//		LockSystem: webdav.NewMemLS(),
		//	},
		//	user,
		//}
		prefix := "/api/v1/webdav" //setting.WebDavPrefix//请求前缀
		rootDir := setting.WebDavDir   //webdir目录

		if requestRoot, err := stripPrefix(prefix, c.Request.URL.Path); err == nil {
			//shareInfo,err = models.ShareInfo(reqPath)
			//relPath := reqPath
			log.Println("reqPath->:", requestRoot)
			//fileSystem := &models.WebDavFs{User:user,Dir:webdav.Dir(dir),Storage:user.Username}
			//fileSystem := &webdav.Dir{
			//	Context:  c.Request.Context(),
			//	RootPath: dir,
			//	User:     user,
			//	//Owner:    user.Username,
			//	BaseReqPath: reqPath,
			//	ReqPath:     reqPath,
			//	//RelPath:  relPath,
			//}
			fileSystem := models.NewWebdavFs(rootDir,prefix,requestRoot,user,c)

			fs := webdav.Handler{
				Prefix:     prefix,
				FileSystem: fileSystem,
				LockSystem: webdav.NewMemLS(),
			}
			//log.Println(fileName)

			//log.Print(fs.FileSystem)
			fs.ServeHTTP(c.Writer, c.Request)
		}
	}
}

func stripPrefix(Prefix, p string) (string, error) {
	if Prefix == "" {
		return p, nil
	}
	if r := strings.TrimPrefix(p, Prefix); len(r) < len(p) {
		return r, nil
	}
	return p, errors.New("地址不对")
}