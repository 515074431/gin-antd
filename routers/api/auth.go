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
	code := e.INVALID_PARAMS

	var form auth
	if c.ShouldBind(&form) == nil {
		log.Println("form", form)
		log.Println("form", form.Username)
		log.Println("form", form.Password)

		username := form.Username
		password := form.Password

		valid := validation.Validation{}
		ok, _ := valid.Valid(&form)

		if ok {

			if isCheck, auth := models.CheckAuth(username, password); isCheck {
				token, err := util.GenerateToken(auth)
				if err != nil {
					code = e.ERROR_AUTH_TOKEN
				} else {
					data["token"] = token
					data["username"] = username
					data["currentAuthority"] = "admin" //临时先设置管理员

					code = e.SUCCESS
				}

			} else {
				data["currentAuthority"] = "guest"
				code = e.ERROR_AUTH
			}
		} else {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

