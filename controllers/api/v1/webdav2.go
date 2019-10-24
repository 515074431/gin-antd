package v1

import (
	"github.com/515074431/gin-antd/pkg/setting"
	"github.com/515074431/gin-antd/pkg/util"
	"github.com/515074431/gin-antd/pkg/webdav"
	"github.com/gin-gonic/gin"

	//"golang.org/x/net/webdav"
	"log"
)

func WebDav2(c *gin.Context) {
	if user, ok := util.GetIdentify(c); ok {
		log.Println("User:", user)
		prefix := "/api/v1/webdav2" //setting.WebDavPrefix//请求前缀
		rootDir := setting.WebDavDir   //webdir目录

		if requestRoot, err := stripPrefix(prefix, c.Request.URL.Path); err == nil {
			log.Println("reqPath->:", requestRoot)
			fileSystem := webdav.NewWebdavFs(rootDir,prefix,requestRoot,user,c)

			fs := webdav.Handler{
				Prefix:     prefix,
				FileSystem: fileSystem,
				LockSystem: webdav.NewMemLS(),
			}
			fs.ServeHTTP(c.Writer, c.Request)
		}
	}
}

