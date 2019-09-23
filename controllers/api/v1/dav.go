package v1

import (
	"github.com/515074431/gin-antd/pkg/webdav"
	"github.com/gin-gonic/gin"
	"log"
)

func Dav(c *gin.Context) {
	//c.Request.Method = "PROPFIND"
	fileName := c.Param("file")

	fs := &webdav.Handler{
		Prefix: "/api/v1/dav/" ,
		FileSystem: webdav.Dir("d:/tmp" ),
		LockSystem: webdav.NewMemLS(),
	}
	log.Println(fileName)

	log.Print(fs.FileSystem)
	fs.ServeHTTP(c.Writer,c.Request)


}


