package user

import (
	"beautyProject/internal/backend/user"
	"beautyProject/internal/pkg/web/gin/handler/response"
	"beautyProject/internal/pkg/web/request"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (h *Handler) Register(c *gin.Context, reqId string, req request.UserReq) {
	userOb := user.User{reqId}
	msgDto := userOb.Register(req)
	response.ProcessMsgDto(c, msgDto)
}

func (h *Handler) Login(c *gin.Context, reqId string, req request.UserReq) {
	userOb := user.User{reqId}
	userOb.Login(req)
	//response.Success(c, "登入成功")
}

func (h *Handler) Logout(c *gin.Context, reqId string, req request.EmptyReq) {
	userOb := user.User{}
	userOb.Logout()
	//response.Success(c, "已登出")
}
