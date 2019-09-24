package util

import (
	"github.com/515074431/gin-antd/models"
	"github.com/515074431/gin-antd/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	UserId int `json:"id"`
	Username string `json:"username"`
	//Password string `json:"password"`
	jwt.StandardClaims
}


func GenerateToken(user models.User) (string, error) {
	nowTime :=time.Now()
	expireTime := nowTime.Add(3* time.Hour)
	claims := jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),// 失效时间
			Audience:  user.Username,     // 受众
			Id:        string(user.ID),   // 编号
			IssuedAt:  time.Now().Unix(), // 签发时间
			Issuer:    "gin-antd",       // 签发人
			NotBefore: time.Now().Unix(), // 生效时间
			Subject:   "login",           // 主题
		}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*jwt.StandardClaims, error) {
	tokenClaims,err := jwt.ParseWithClaims(token,&jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret,nil
	})
	if tokenClaims !=nil {
		if claims, ok := tokenClaims.Claims.(*jwt.StandardClaims); ok && tokenClaims.Valid {
			return claims,nil
		}
	}
	return nil,err
}

func GetIdentify(c *gin.Context) (user *models.User,err error) {
//func GetIdentify(c *gin.Context) (identify *Claims,err error) {

	identify := c.MustGet("Identify").(*jwt.StandardClaims)

	//user.ID = com.StrTo(identify.UserId).MustInt()
	err =models.Db.First(&user,identify.Id).Error
	return
}
