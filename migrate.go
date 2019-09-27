package main

import (
	"github.com/515074431/gin-antd/models"
	"log"
	"os"
)

func main()  {
	if !models.Db.HasTable(&models.User{}){
		models.Db.CreateTable(&models.User{})
		log.Println("成功创建表：",models.User{}.TableName())
	}
	if !models.Db.HasTable(&models.Article{}){
		models.Db.CreateTable(&models.Article{})
		log.Println("成功创建表：Article")
	}
	if !models.Db.HasTable(&models.Tag{}){
		models.Db.CreateTable(&models.Tag{})
		log.Println("成功创建表：Tag")
	}
	if !models.Db.HasTable(&models.Auth{}){
		models.Db.CreateTable(&models.Auth{})
		log.Println("成功创建表：Auth")
	}
	if !models.Db.HasTable(&models.Share{}){
		models.Db.CreateTable(&models.Share{})
		log.Println("成功创建表：Share")
	}

	os.Exit(0)




	models.CloseDB()
}
