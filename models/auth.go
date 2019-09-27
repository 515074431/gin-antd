package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Auth struct {
	gorm.Model
	Username string `gorm:"username" json:"username"`
	Password string `gorm:"password" json:"password"`
}
//var UserInfo *Auth
//检查权限
func CheckAuth(username,password string) (err error,user User) {
	err,user = UserFindByName(username)
	if  err == nil{
		if user.Password != password{
			err = errors.New("密码错误")
		}
	}else{
		err = errors.New("账户不存在")
	}
	return
}

//func (auth *Auth) UserInfo()  *Auth {
//	return UserInfo
//}