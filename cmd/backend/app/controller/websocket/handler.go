package websocket

import (
	"beautyProject/internal/pkg/enum/mouseAction"
	"beautyProject/internal/pkg/interfaces"
	commonRecord "beautyProject/internal/services/backend/common/record"
	"beautyProject/internal/services/backend/user/mouseAction/record"
)

type Handler struct {
}

func (handler *Handler) RecordAction(httpInfo *HttpInfo, msg []byte) {
	userId := httpInfo.GetUserID()
	c := httpInfo.GetGinContext()
	actionStr := c.DefaultQuery("action", "unknown") //雖然等於連線管理參數，但意義不太依樣
	action := handler.getAction(actionStr)
	action.Work(userId, msg)
}

func (handler *Handler) getAction(action string) interfaces.IWork {
	switch action {
	case mouseAction.Move.ChineseName:
		return &record.Move{Recorder: commonRecord.Recorder{Name: mouseAction.Move.Name}}
	case mouseAction.Click.ChineseName:
		return &record.Click{Recorder: commonRecord.Recorder{Name: mouseAction.Click.Name}}
	case mouseAction.Scroll.ChineseName:
		return &record.Scroll{Recorder: commonRecord.Recorder{Name: mouseAction.Scroll.Name}}
	}
	return nil
}
