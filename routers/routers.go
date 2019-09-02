package routers
//路由处理中心
import (
	"github.com/515074431/gin-antd/middleware/jwt"
	"github.com/515074431/gin-antd/pkg/setting"
	"github.com/515074431/gin-antd/routers/api"
	"github.com/515074431/gin-antd/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

var identityKey = "id"
//设置路由
func SetupRouter() *gin.Engine {
	r := gin.New()










	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)


	r.StaticFS("/public",http.Dir("public"))

	//添加Get路由
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello gin")
	})
	//获取权限
	r.GET("/auth",api.GetAuth)

	//r.POST("/login", authMiddleware.LoginHandler)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id",v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id",v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}