package record

import (
	"beautyProject/internal/pkg/enum/authLocation"
	"beautyProject/internal/pkg/util/web/gin/middleware/router"
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	handler := Handler{}
	rg.GET("/average-info",
		router.Integrate(handler.AverageRecordInfo, router.Auth{IsAuth: true, AuthLocation: authLocation.Cookie}))
	rg.GET("/raw-info",
		router.Integrate(handler.RawRecordInfo, router.Auth{IsAuth: true, AuthLocation: authLocation.Cookie}))
}
