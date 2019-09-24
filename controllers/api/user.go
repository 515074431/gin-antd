package api

import (
	"github.com/515074431/gin-antd/models"
	"github.com/515074431/gin-antd/pkg/e"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	//msg := make(map[string]interface{})
	register := &models.Register{}
	result := e.NewDefaultResult()

	if err := c.ShouldBind(&register); err != nil {
		result.Message = err.Error()
		result.Code = e.ERROR
		c.JSON(result.Code,result)
		return
	}

	if user, err := register.Save(); err == nil {
		result.Status = true
		result.Data = map[string]interface{}{
			"id":       user.ID,
			"email":    user.Email,
			"username": user.Username,
			"role":     user.Role,
			"err":      err,}


	} else {

		result.Code = e.ERROR
		result.Message = err.Error()
	}
	c.JSON(result.Code, result)

}
