package v1

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/webdav"
	"log"
)

type Webdav  struct {
	fs webdav.Handler

}

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


