package util

import (
	"github.com/515074431/gin-antd/models"
	"github.com/515074431/gin-antd/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	UserId   int    `json:"id"`
	Username string `json:"username"`
	//Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(user models.User) (string, error) {
	nowTime := time.Now()
	tokenExpireHour :=setting.TokenExpireHour
	expireTime := nowTime.Add(time.Duration(tokenExpireHour) * time.Hour)
	//expireTime := nowTime.Add(setting.TokenExpireHour * time.Hour)
	claims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),                  // 失效时间
		Audience:  user.Username,                      // 受众
		Id:        string(strconv.Itoa(int(user.ID))), // 编号
		IssuedAt:  time.Now().Unix(),                  // 签发时间
		Issuer:    "gin-antd",                         // 签发人
		NotBefore: time.Now().Unix(),                  // 生效时间
		Subject:   "login",                            // 主题
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*jwt.StandardClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.StandardClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func GetIdentify(c *gin.Context) (user models.User, ok bool) {
	user, ok = c.MustGet("Identify").(models.User)
	return
}
