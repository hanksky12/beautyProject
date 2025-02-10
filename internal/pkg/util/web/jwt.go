package web

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	env "github.com/spf13/viper"
	"strconv"
	"time"
)

type Claims struct {
	UserId uint `json:"user_id"`
	jwt.RegisteredClaims
}

const CookieName = "token"

type Jwt struct {
	GinContext *gin.Context
}

func (j *Jwt) SetCookie(userId uint) error {
	log.Info("SetJwtCookie Start")
	var envJwt = env.GetStringMapString("Jwt")
	minutes, err := strconv.Atoi(envJwt["access_token_expires"])
	if err != nil {
		return err
	}
	expirationTime := time.Now().Add(time.Minute * time.Duration(minutes))
	claims := Claims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),     // 生效时间
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 使用HS256签名算法
	tokenString, err := t.SignedString([]byte(envJwt["secret_key"]))
	if err != nil {
		return err
	}
	j.GinContext.SetCookie(CookieName, tokenString, int(expirationTime.Sub(time.Now()).Seconds()), "/", "localhost", false, true)
	log.Info("SetJwtCookie Success")
	return nil
}

func (j *Jwt) UnsetCookie() {
	j.GinContext.SetCookie(CookieName, "", -1, "/", "localhost", false, true)
}

func (j *Jwt) GetToken() (string, error) {
	tokenString, err := j.GinContext.Cookie(CookieName)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *Jwt) Parse(tokenStr string) (*Claims, error) {
	var envJwt = env.GetStringMapString("Jwt")
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(envJwt["secret_key"]), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
