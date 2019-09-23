package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password" gorm:"type:varchar(100);"`

	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"` // 设置字段大小为255
	//MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
}
// 通过TableName方法将User表命名为`profiles`
func (User) TableName() string {
	return "profiles"
}
//注册模型
type Register struct {
	Email string ` form:"email" json:"email"  binding:"email"`
	Username string `form:"username"  json:"username"  binding:"required"`
	Password string ` form:"password" json:"password"  binding:"required"`
	PasswordAgain string `form:"password-again" json:"password-again"  binding:"eqfield=Password"`
}
//注册时保存用户
func (r Register) Save() ( uer User ,err error ){

	user := User{
		Username:r.Username,
		Password: r.Password,//密码后面需要加密处理
		Email:r.Email,
		Role:"user",
	}
	if errors :=Db.Create(&user); errors != nil{
		err = errors.Error
	}

	return user,err
}
//登录模型
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
//检查权限
func CheckUser(username, password string) (isCheck bool,UserInfo User) {
	Db.Select("id").Where(Auth{Username: username, Password: password}).First(&UserInfo)
	if UserInfo.ID > 0 {
		isCheck = true
		}
	return
}

//func (auth *Auth) UserInfo()  *Auth {
//	return UserInfo
//}