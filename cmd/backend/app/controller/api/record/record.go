package record

import (
	"beautyProject/internal/pkg/util/web/gin/middleware/router"
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	handler := Handler{}
	rg.GET("/average-info", router.Integrate(handler.AverageRecordInfo, true))
	rg.GET("/raw-info", router.Integrate(handler.RawRecordInfo, true)) // 新增原始紀錄路由
}
