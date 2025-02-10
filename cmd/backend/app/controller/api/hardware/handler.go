package hardware

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/enum"
	"beautyProject/internal/pkg/interfaces"
	"beautyProject/internal/pkg/util/web/gin/handler/response"
	"beautyProject/internal/pkg/web/request"
	"beautyProject/internal/services/backend/hardware/info"
	"beautyProject/internal/services/backend/hardware/record"
	"beautyProject/internal/services/backend/hardware/record/base"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler struct{}

func (handler *Handler) RecordHardware(c *gin.Context, req request.HardwareReq) {
	var msgDto dto.Msg

	hardware := handler.getHardware(req)
	userId := strconv.FormatUint(uint64(c.GetUint("userId")), 10)

	recorder := base.NewRecorder(hardware, userId)
	recorder.Initialize()
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

func (handler *Handler) getHardware(req request.HardwareReq) interfaces.IWorkAndName {
	switch req.Hardware {
	case enum.Cpu.Name:
		return &record.Cpu{}
	case enum.Disk.Name:
		return &record.Disk{}
	case enum.Memory.Name:
		return &record.Memory{}
	}
	return nil
}

func (handler *Handler) HardwareInfo(c *gin.Context, req request.EmptyReq) {
	hardware := &info.HardwareInfo{}
	tableDto := hardware.Query()
	response.ProcessTableDto(c, tableDto)
}
