package web

import (
	"beautyProject/internal/pkg/enum/authLocation"
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

const TokenName = "token"

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
	j.GinContext.SetCookie(TokenName, tokenString, int(expirationTime.Sub(time.Now()).Seconds()), "/", "localhost", false, true)
	log.Info("SetJwtCookie Success")
	return nil
}

func (j *Jwt) UnsetCookie() {
	j.GinContext.SetCookie(TokenName, "", -1, "/", "localhost", false, true)
}

func (j *Jwt) GetToken(authLoc *authLocation.AuthLocation) (string, error) {
	switch authLoc {
	case authLocation.Cookie:
		tokenString, err := j.GinContext.Cookie(TokenName)
		if err != nil {
			return "", err
		}
		return tokenString, nil
	case authLocation.QueryParams:
		tokenString := j.GinContext.DefaultQuery(TokenName, "unknown")
		if tokenString == "unknown" {
			return "", nil
		}
		return tokenString, nil
	default:
		return "", nil
	}
}

func (j *Jwt) Parse(tokenStr string) (*Claims, error) {
	var envJwt = env.GetStringMapString("Jwt")
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(envJwt["secret_key"]), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
