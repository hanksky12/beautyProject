package hardware

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/enum"
	"beautyProject/internal/pkg/util/web/gin/handler/response"
	"beautyProject/internal/pkg/web/request"
	"beautyProject/internal/services/backend/pc"
	"beautyProject/internal/services/backend/pc/base"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler struct{}

func (h *Handler) RecordHardware(c *gin.Context, req request.HardwareReq) {
	userId := strconv.FormatUint(uint64(c.GetUint("userId")), 10)
	recorder := base.NewRecorder(h.getHardware(req), userId)
	recorder.Initialize()
	var msgDto dto.Msg
	switch req.State {
	case enum.Start.Name:
		msgDto = recorder.Start()
	case enum.Stop.Name:
		msgDto = recorder.Stop()
	default:
		msgDto = dto.Msg{Success: false, Message: "未知操作"}
	}
	response.ProcessMsgDto(c, msgDto)
}

func (h *Handler) getHardware(req request.HardwareReq) base.WorkAndName {
	switch req.Hardware {
	case enum.Cpu.Name:
		return &pc.Cpu{}
	case enum.Disk.Name:
		return &pc.Disk{}
	case enum.Memory.Name:
		return &pc.Memory{}
	}
	return nil
}
