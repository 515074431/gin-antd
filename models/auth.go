package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `gorm:"username" json:"username"`
	Password string `gorm:"password" json:"password"`
}
//var UserInfo *Auth
//检查权限
func CheckAuth(username, password string) (isCheck bool,UserInfo Auth) {
	db.Select("*").Where(Auth{Username: username, Password: password}).First(&UserInfo)
	if UserInfo.ID > 0 {
		isCheck = true
		}
	return
}

//func (auth *Auth) UserInfo()  *Auth {
//	return UserInfo
//}