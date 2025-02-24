package websocket

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type HttpInfo struct {
	C                      *gin.Context
	route                  string
	userID                 string
	connectionsManageParam string
}

func (handler *HttpInfo) Analyze(connectionsManageParam string) {
	handler.route = handler.C.FullPath()
	handler.userID = strconv.FormatUint(uint64(handler.C.GetUint("userId")), 10)
	handler.connectionsManageParam = handler.C.DefaultQuery(connectionsManageParam, "unknown")
}
func (handler *HttpInfo) GetRoute() string {
	return handler.route
}
func (handler *HttpInfo) GetUserID() string {
	return handler.userID
}
func (handler *HttpInfo) GetConnectionsManageParam() string {
	return handler.connectionsManageParam
}
func (handler *HttpInfo) GetGinContext() *gin.Context {
	return handler.C
}
