package user

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/repository"
	"beautyProject/internal/pkg/util/web/gin/handler/response"
	"beautyProject/internal/pkg/web/request"
	"beautyProject/internal/services/backend/user"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler struct{}

func (h *Handler) Register(c *gin.Context, req request.UserReq) {
	userRepo := repository.User{}
	userServ := user.User{}
	msgDto := userServ.Register(c, userRepo, req)
	response.ProcessMsgDto(c, msgDto)
}

func (h *Handler) Login(c *gin.Context, req request.UserReq) {
	userRepo := repository.User{}
	userServ := user.User{}
	msgDto := userServ.Login(c, userRepo, req)
	response.ProcessMsgDto(c, msgDto)
}

func (h *Handler) Logout(c *gin.Context, req request.EmptyReq) {
	userOb := user.User{}
	msgDto := userOb.Logout(c)
	response.ProcessMsgDto(c, msgDto)
}

func (h *Handler) Info(c *gin.Context, req request.EmptyReq) {
	userId := strconv.FormatUint(uint64(c.GetUint("userId")), 10)
	msgDto := dto.Msg{Success: true, Message: userId}
	response.ProcessMsgDto(c, msgDto)
}
