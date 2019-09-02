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
	UserId int `json:"user_id"`
	Username string `json:"username"`
	//Password string `json:"password"`
	jwt.StandardClaims
}


func GenerateToken(auth models.Auth) (string, error) {
	nowTime :=time.Now()
	expireTime := nowTime.Add(3* time.Hour)
	claims := Claims{
		auth.ID,
		auth.Username,
		//auth.Password,
		jwt.StandardClaims{
			ExpiresAt:expireTime.Unix(),
			Issuer:"gin-antd",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims,err := jwt.ParseWithClaims(token,&Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret,nil
	})
	if tokenClaims !=nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims,nil
		}
	}
	return nil,err
}

func GetIdentify(c *gin.Context) (user *models.Auth,err error) {
//func GetIdentify(c *gin.Context) (identify *Claims,err error) {
	user = &models.Auth{}
	identify := c.MustGet("Identify").(*Claims)

	//user.ID = com.StrTo(identify.UserId).MustInt()
	user.ID = identify.UserId
	user.Username = identify.Username

	return
}
