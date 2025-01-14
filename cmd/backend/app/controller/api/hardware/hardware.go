package hardware

import (
	"beautyProject/internal/pkg/util/web/gin/middleware/router"
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	handler := Handler{}
	rg.POST("/status", router.Integrate(handler.RecordHardware, true))
	rg.GET("/info", router.Integrate(handler.HardwareInfo, true))

}
