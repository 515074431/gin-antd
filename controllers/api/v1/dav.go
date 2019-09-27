package v1

import (
	"github.com/515074431/gin-antd/models"
	"github.com/515074431/gin-antd/pkg/setting"
	"github.com/515074431/gin-antd/pkg/util"
	//"github.com/515074431/gin-antd/pkg/webdav"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/webdav"
	"log"
)

func Dav(c *gin.Context) {
	//c.Request.Method = "PROPFIND"
	fileName := c.Param("file")

	if user, ok := util.GetIdentify(c); ok {
		log.Println("User:", user)
		//fs := &webdav.Server{
		//	webdav.Handler{Prefix: "/api/v1/dav/",
		//		FileSystem: webdav.Dir("d:/tmp/"),
		//		LockSystem: webdav.NewMemLS(),
		//	},
		//	user,
		//}
		prefix:=setting.WebDavPrefix//请求前缀
		dir := setting.WebDavDir//webdir目录
		fileSystem := &models.WebDavFs{User:user,Dir:webdav.Dir(dir),Storage:user.Username}
		fileSystem.SetStartReqName(prefix,c.Request.URL.Path)//一定要先执行一下这个

		fs := &webdav.Handler{
			Prefix:prefix,
			FileSystem: fileSystem,
			LockSystem:webdav.NewMemLS(),
		}
		log.Println(fileName)

		log.Print(fs.FileSystem)
		fs.ServeHTTP(c.Writer, c.Request)
	}

}