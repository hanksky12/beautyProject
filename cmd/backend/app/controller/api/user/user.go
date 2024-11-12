package user

import (
	"beautyProject/internal/pkg/util/web/gin/middleware/router"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	handler := Handler{}
	rg.POST("/user",
		router.Integrate(handler.Register, false))
	rg.POST("/user/login",
		router.Integrate(handler.Login, false))
	rg.GET("/user/logout",
		router.Integrate(handler.Logout, false))
	rg.GET("/user/protect_info",
		router.Integrate(handler.Info, true))
}
