package jwt

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/515074431/gin-antd/models"
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
			code = e.ERROR_UNAUTHORIZED
		} else {
			claims, err := util.ParseToken(token)

			if err != nil {
				code = e.ERROR_FORBIDDEN
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_FORBIDDEN
			}else {
				if err,user :=models.UserFindByName(claims.Audience); err == nil  {
					log.Println("user:",user)
					if user.UpdatedAt.Unix() > claims.IssuedAt{ //更新时间大于生效时间，需要重新登录
						code = e.ERROR_FORBIDDEN
					}else{
						c.Set("Identify", claims)
						log.Println("claims: ",claims)
					}
				}else{
					code = e.ERROR_UNAUTHORIZED
				}
			}


		}
		if code != e.SUCCESS {
			result := e.Result{
				Code:code,
				Message:e.GetMsg(code),
			}
			c.JSON(http.StatusUnauthorized, result)

			c.Abort()
			return
		}


		c.Next()

	}
}

