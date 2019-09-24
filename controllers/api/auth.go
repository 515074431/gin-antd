package api

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/515074431/gin-antd/models"
	"github.com/515074431/gin-antd/pkg/e"
	"github.com/515074431/gin-antd/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}



func GetAuth(c *gin.Context) {
	data := make(map[string]interface{})
	result := e.Result{}

	var form auth
	if c.ShouldBind(&form) == nil {
		username := form.Username
		password := form.Password

		valid := validation.Validation{}
		ok, _ := valid.Valid(&form)

		if ok {

			if err, user := models.CheckAuth(username,password); err == nil {
				token, err := util.GenerateToken(user)
				if err != nil {
					result.Code = e.ERROR_UNAUTHORIZED
				} else {
					data["token"] = token
					data["username"] = username
					data["currentAuthority"] = "admin" //临时先设置管理员

					result.Code = e.SUCCESS
					result.Data = data
				}

			} else {
				data["currentAuthority"] = "guest"
				result.Code = e.ERROR_FORBIDDEN
				result.Data = data
			}
		} else {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			result.Error = valid.Errors
		}
	}

	c.JSON(http.StatusOK, result)
}

