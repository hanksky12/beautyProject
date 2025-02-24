package authentication

import (
	"beautyProject/internal/pkg/enum/authLocation"
	"beautyProject/internal/pkg/util/web/gin/middleware/router"
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	handler := Handler{}
	rg.POST("/register",
		router.Integrate(handler.Register, router.Auth{IsAuth: false, AuthLocation: authLocation.Cookie}))
	rg.POST("/login",
		router.Integrate(handler.Login, router.Auth{IsAuth: false, AuthLocation: authLocation.Cookie}))
	rg.GET("/logout",
		router.Integrate(handler.Logout, router.Auth{IsAuth: false, AuthLocation: authLocation.Cookie}))
	rg.GET("/token",
		router.Integrate(handler.Token, router.Auth{IsAuth: true, AuthLocation: authLocation.Cookie}))
	rg.GET("/protect_info",
		router.Integrate(handler.Info, router.Auth{IsAuth: true, AuthLocation: authLocation.Cookie}))
}
