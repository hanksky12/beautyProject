package hardware

import (
	"beautyProject/internal/pkg/dto"
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
	case "start":
		msgDto = recorder.Start()
	case "stop":
		msgDto = recorder.Stop()
	}
	response.ProcessMsgDto(c, msgDto)
}

func (h *Handler) getHardware(req request.HardwareReq) base.WorkAndName {
	switch req.Hardware {
	case "cpu":
		return &pc.Cpu{}
	case "memory":
		return &pc.Memory{}
	case "disk":
		return &pc.Disk{}
	}
	return nil
}
