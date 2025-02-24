package router

import (
	"beautyProject/internal/pkg/enum/authLocation"
	"beautyProject/internal/pkg/util/web"
	"beautyProject/internal/pkg/util/web/gin/handler/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Auth struct {
	IsAuth       bool
	AuthLocation *authLocation.AuthLocation
}

func authentication(c *gin.Context, auth Auth) (bool, string) {
	if !auth.IsAuth {
		return true, ""
	}
	jwt := web.Jwt{GinContext: c}
	tokenString, err := jwt.GetToken(auth.AuthLocation)
	if err != nil {
		response.Unauthorized(c)
		c.Abort()
		return false, ""
	}
	claims, err := jwt.Parse(tokenString)
	if err != nil {
		response.Unauthorized(c)
		c.Abort()
		return false, ""
	}
	c.Set("userId", claims.UserId)
	return true, strconv.Itoa(int(claims.UserId))

}
