package models

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	FirstName string
	LastName  string
}
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
//检查权限
func CheckUser(username, password string) (isCheck bool,UserInfo User) {
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&UserInfo)
	if UserInfo.ID > 0 {
		isCheck = true
		}
	return
}

//func (auth *Auth) UserInfo()  *Auth {
//	return UserInfo
//}