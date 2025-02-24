package hardwareSpecs

import (
	"beautyProject/internal/pkg/enum/authLocation"
	"beautyProject/internal/pkg/util/web/gin/middleware/router"
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	handler := Handler{}
	rg.GET("/info",
		router.Integrate(handler.HardwareSpecsInfo, router.Auth{IsAuth: true, AuthLocation: authLocation.Cookie}))

}
