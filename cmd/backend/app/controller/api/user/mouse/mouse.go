package mouse

import (
	"beautyProject/internal/pkg/enum/authLocation"
	"beautyProject/internal/pkg/util/web/gin/middleware/router"
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	handler := Handler{}
	rg.GET("/tracking",
		router.Integrate(handler.RecordAction, router.Auth{IsAuth: true, AuthLocation: authLocation.QueryParams}))
	rg.GET("/info",
		router.Integrate(handler.ActionInfo, router.Auth{IsAuth: true, AuthLocation: authLocation.Cookie}))

}
