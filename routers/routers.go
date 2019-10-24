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
	r.GET("/api/auth-routes", func(context *gin.Context) {
		//authRoutes := `{"/":{"authority":["admin","user"]}}`
		//str := `{"/":{"authority":["admin","user"]}}`
		//authRoutes :=
		//json.Unmarshal([]byte(str),authRoutes)
		authRoutes := map[string]map[string][]string{"/":{"authority":{"admin","user"}}}
		context.JSON(http.StatusOK,authRoutes)
	})
	r.GET("/api/current-user", func(context *gin.Context) {
		//currentUser := map[string]interface{}
		str := `{"name":"David Fang","avatar":"https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png","userid":"00000001","email":"fangmw@lichengsoft.com","signature":"海纳百川，有容乃大","title":"交互专家","group":"励铖软件－建筑云盘技术部－研发工程师","tags":[{"key":"0","label":"很有想法的"},{"key":"1","label":"专注设计"},{"key":"2","label":"辣~"},{"key":"3","label":"大长腿"},{"key":"4","label":"川妹子"},{"key":"5","label":"海纳百川"}],"notifyCount":12,"unreadCount":11,"country":"China","geographic":{"province":{"label":"上海市","key":"310000"},"city":{"label":"静安区","key":"310100"}},"address":"新闸路848号6楼","phone":"021-268888888"}`
		context.String(http.StatusOK,str)

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
		apiv1.Handle("PROPFIND", "/webdav/*file", apiV1.WebDav)
		apiv1.Handle("PROPFIND", "/webdav2/*file", apiV1.WebDav2)

		apiv1.Handle("PUT", "/webdav/*file", apiV1.WebDav)
		apiv1.Handle("DELETE", "/webdav/*file", apiV1.WebDav)
		apiv1.Handle("POST", "/webdav/*file", apiV1.WebDav)
		apiv1.Handle("MOVE", "/webdav/*file", apiV1.WebDav)
		apiv1.Handle("COPY", "/webdav/*file", apiV1.WebDav)
		apiv1.Handle("MKCOL", "/webdav/*file", apiV1.WebDav)
		apiv1.Handle("LOCK", "/webdav/*file", apiV1.WebDav)
		apiv1.Handle("UNLOCK", "/webdav/*file", apiV1.WebDav)
		apiv1.Handle("PROPPATCH", "/webdav/*file", apiV1.WebDav)
		apiv1.Handle("OPTIONS", "/webdav/*file", apiV1.WebDav)

		apiv1.Handle("PUT", "/webdav2/*file", apiV1.WebDav2)
		apiv1.Handle("DELETE", "/webdav2/*file", apiV1.WebDav2)
		apiv1.Handle("POST", "/webdav2/*file", apiV1.WebDav2)
		apiv1.Handle("MOVE", "/webdav2/*file", apiV1.WebDav2)
		apiv1.Handle("COPY", "/webdav2/*file", apiV1.WebDav2)
		apiv1.Handle("MKCOL", "/webdav2/*file", apiV1.WebDav2)
		apiv1.Handle("LOCK", "/webdav2/*file", apiV1.WebDav2)
		apiv1.Handle("UNLOCK", "/webdav2/*file", apiV1.WebDav2)
		apiv1.Handle("PROPPATCH", "/webdav2/*file", apiV1.WebDav2)
		apiv1.Handle("OPTIONS", "/webdav2/*file", apiV1.WebDav2)

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