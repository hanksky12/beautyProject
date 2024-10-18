package user

import (
	"beautyProject/internal/pkg/web/gin/handler/response"
	"beautyProject/internal/pkg/web/request"
	"beautyProject/internal/services/backend/user"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (h *Handler) Register(c *gin.Context, req request.UserReq) {
	userServ := user.User{}
	msgDto := userServ.Register(req)
	response.ProcessMsgDto(c, msgDto)
}

func (h *Handler) Login(c *gin.Context, req request.UserReq) {
	userOb := user.User{}
	userOb.Login(req)
	//response.Success(c, "登入成功")
}

func (h *Handler) Logout(c *gin.Context, req request.EmptyReq) {
	userOb := user.User{}
	userOb.Logout()
	//response.Success(c, "已登出")
}
