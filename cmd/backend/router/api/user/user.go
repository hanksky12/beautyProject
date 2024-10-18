package user

import (
	"beautyProject/internal/pkg/web/gin/middleware/router"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	handler := Handler{}
	rg.POST("/user", router.Integrate(handler.Register))
	rg.POST("/user/login", router.Integrate(handler.Login))
	rg.GET("/user/logout", router.Integrate(handler.Logout))
}
