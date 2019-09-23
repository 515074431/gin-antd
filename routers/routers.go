package routers
//路由处理中心
import (
	"github.com/515074431/gin-antd/controllers/api"
	apiV1 "github.com/515074431/gin-antd/controllers/api/v1"
	"github.com/515074431/gin-antd/middleware/jwt"
	"github.com/515074431/gin-antd/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

var identityKey = "id"
//设置路由
func SetupRouter() *gin.Engine {
	//r := gin.New()
	r := gin.Default()

	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)



	r.StaticFS("/public",http.Dir("public"))

	//添加Get路由
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello gin")
	})
	//用户注册
	r.POST("/api/v1/register",api.UserRegister)
	//获取权限
	r.POST("/api/v1/account", api.GetAuth)

	//r.POST("/login", authMiddleware.LoginHandler)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", apiV1.GetTags)
		//新建标签
		apiv1.POST("/tags", apiV1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", apiV1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", apiV1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", apiV1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", apiV1.GetArticle)
		//新建文章
		apiv1.POST("/articles", apiV1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", apiV1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", apiV1.DeleteArticle)

		//WEBDAV方法
		apiv1.Handle("PROPFIND", "/dav/*file", apiV1.Dav)
		//apiv1.Handle("PUR", "/dav/*file", v1.Dav)
		apiv1.Handle("COPY", "/dav/*file", apiV1.Dav)
		apiv1.Handle("MOVE", "/dav/*file", apiV1.Dav)
		apiv1.Handle("PUT", "/dav/*file", apiV1.Dav)
		apiv1.Handle("DELETE", "/dav/*file", apiV1.Dav)
		apiv1.Handle("POST", "/dav/*file", apiV1.Dav)
		apiv1.Handle("MKCOL", "/dav/*file", apiV1.Dav)
		apiv1.Handle("LOCK", "/dav/*file", apiV1.Dav)
		apiv1.Handle("UNLOCK", "/dav/*file", apiV1.Dav)
		apiv1.Handle("PROPPATCH", "/dav/*file", apiV1.Dav)
		apiv1.Handle("OPTIONS", "/dav/*file", apiV1.Dav)

		//apiv1.Any("/dav/*file", v1.Dav)
	}
	//r.Handle("PROPFIND", "/api/v1/dav/", v1.Dav)
	//r.Handle("COPY", "/dav/", v1.Dav)
	return r
}