package jwt

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/515074431/gin-antd/pkg/e"
	"github.com/515074431/gin-antd/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		//var data interface{}

		code = e.SUCCESS
		//token := c.Query("token")
		token :=  c.Request.Header.Get("Authorization")

		if s := strings.Split(token, " "); len(s) == 2 {
			token = s[1]
		}

		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			// 继续交由下一个路由处理,并将解析出的信息传递下去
			c.Set("Identify", claims)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				//"data": data,
			})

			c.Abort()
			return
		}


		c.Next()

	}
}

