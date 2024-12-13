package router

import (
	"beautyProject/internal/pkg/util/web"
	"beautyProject/internal/pkg/util/web/gin/handler/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func auth(c *gin.Context, isAuth bool) (bool, string) {
	if !isAuth {
		return true, ""
	}
	tokenString, err := web.GetJwtToken(c)
	if err != nil {
		response.Unauthorized(c)
		c.Abort()
		return false, ""
	}
	claims, err := web.ParseJwt(tokenString)
	if err != nil {
		response.Unauthorized(c)
		c.Abort()
		return false, ""
	}
	c.Set("userId", claims.UserId)
	return true, strconv.Itoa(int(claims.UserId))

}
