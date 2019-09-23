package main

import (
	"github.com/515074431/gin-antd/models"
)

func main()  {
	models.Db.AutoMigrate(&models.User{})
	models.Db.AutoMigrate(&models.Article{})
	models.Db.AutoMigrate(&models.Tag{})
	models.Db.AutoMigrate(&models.Auth{})
	models.CloseDB()
}
