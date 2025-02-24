package hardwareSpecs

import (
	"beautyProject/internal/pkg/util/web/gin/handler/response"
	"beautyProject/internal/pkg/web/request"
	"beautyProject/internal/services/backend/server/hardwareSpecs/info"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (handler *Handler) HardwareSpecsInfo(c *gin.Context, req request.EmptyReq) {
	hardware := &info.HardwareSpecsInfo{}
	tableDto := hardware.Query()
	response.ProcessTableDto(c, tableDto)
}
