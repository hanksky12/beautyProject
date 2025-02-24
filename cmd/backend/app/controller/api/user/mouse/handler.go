package mouse

import (
	"beautyProject/cmd/backend/app/controller/websocket"
	"beautyProject/internal/pkg/util/web/gin/handler/response"
	"beautyProject/internal/pkg/web/request"
	"beautyProject/internal/services/backend/user/mouseAction/info"
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
}

func (handler *Handler) RecordAction(c *gin.Context, req request.MouseActionStateReq) {
	httpInfo := &websocket.HttpInfo{C: c}
	httpInfo.Analyze("action")
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "httpInfo", httpInfo)
	if err := websocket.WebSocket.HandleRequest(c.Writer, c.Request.WithContext(ctx)); err != nil {
		log.Infof("Failed to handle WebSocket request: %v", err)
	}
}

func (handler *Handler) ActionInfo(c *gin.Context, req request.EmptyReq) {
	act := &info.ActionInfo{}
	tableDto := act.Query()
	response.ProcessTableDto(c, tableDto)
}
