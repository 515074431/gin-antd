package main

import (
	"fmt"
	"github.com/515074431/gin-antd/pkg/setting"
	"net/http"

	"github.com/515074431/gin-antd/routers"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

//路由开始
func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		//c.String(http.StatusOK, "pong")
		c.JSON(200, gin.H{
			"message":"pong",
		})
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}
//主程序
func main() {
	r := routers.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%d", setting.HTTPPort))
	//_ = r.Run()

	//s := &http.Server{
	//	Addr:fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:r,
	//	ReadTimeout:setting.ReadTimeout,
	//	WriteTimeout:setting.WriteTimeout,
	//	MaxHeaderBytes: 1<<20,
	//}
	//
	//s.ListenAndServe()
}
