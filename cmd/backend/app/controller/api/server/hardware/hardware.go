package hardware

import (
	"beautyProject/internal/pkg/enum/authLocation"
	"beautyProject/internal/pkg/util/web/gin/middleware/router"
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	handler := Handler{}
	rg.POST("/status",
		router.Integrate(handler.RecordHardware, router.Auth{IsAuth: true, AuthLocation: authLocation.Cookie}))
	rg.GET("/info",
		router.Integrate(handler.HardwareInfo, router.Auth{IsAuth: true, AuthLocation: authLocation.Cookie}))

}
