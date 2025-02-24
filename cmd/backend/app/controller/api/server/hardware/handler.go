package hardware

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/enum/action"
	enumHardware "beautyProject/internal/pkg/enum/hardware"
	"beautyProject/internal/pkg/interfaces"
	"beautyProject/internal/pkg/util/web/gin/handler/response"
	"beautyProject/internal/pkg/web/request"
	commonRecord "beautyProject/internal/services/backend/common/record"
	"beautyProject/internal/services/backend/server/hardware/info"
	"beautyProject/internal/services/backend/server/hardware/record"
	"beautyProject/internal/services/backend/server/hardware/record/base"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler struct{}

func (handler *Handler) RecordHardware(c *gin.Context, req request.HardwareStateReq) {
	var msgDto dto.Msg

	hardware := handler.getHardware(req)
	userId := strconv.FormatUint(uint64(c.GetUint("userId")), 10)

	recorder := base.NewRecorder(hardware, userId)
	recorder.Initialize()
	switch req.State {
	case action.Start.Name:
		msgDto = recorder.Start()
	case action.Stop.Name:
		msgDto = recorder.Stop()
	default:
		msgDto = dto.Msg{Success: false, Message: "未知操作"}
	}
	response.ProcessMsgDto(c, msgDto)
}

func (handler *Handler) getHardware(req request.HardwareStateReq) interfaces.IWorkAndName {
	switch req.Hardware {
	case enumHardware.Cpu.ChineseName:
		return &record.Cpu{Hardware: commonRecord.Recorder{Name: enumHardware.Cpu.Name}}
	case enumHardware.Disk.ChineseName:
		return &record.Disk{Hardware: commonRecord.Recorder{Name: enumHardware.Disk.Name}}
	case enumHardware.Memory.ChineseName:
		return &record.Memory{Hardware: commonRecord.Recorder{Name: enumHardware.Memory.Name}}
	}
	return nil
}

func (handler *Handler) HardwareInfo(c *gin.Context, req request.EmptyReq) {
	hardware := &info.HardwareInfo{}
	tableDto := hardware.Query()
	response.ProcessTableDto(c, tableDto)
}
