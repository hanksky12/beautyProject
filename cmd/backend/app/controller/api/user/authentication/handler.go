package authentication

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/repository"
	"beautyProject/internal/pkg/util/web"
	"beautyProject/internal/pkg/util/web/gin/handler/response"
	"beautyProject/internal/pkg/web/request"
	"beautyProject/internal/services/backend/user/authentication"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler struct{}

func (h *Handler) Register(c *gin.Context, req request.UserReq) {
	userRepo := &repository.User{}
	userServ := authentication.User{}
	jwt := &web.Jwt{GinContext: c}
	msgDto := userServ.Register(jwt, userRepo, req)
	response.ProcessMsgDto(c, msgDto)
}

func (h *Handler) Login(c *gin.Context, req request.UserReq) {
	userRepo := &repository.User{}
	userServ := authentication.User{}
	jwt := &web.Jwt{GinContext: c}
	msgDto := userServ.Login(jwt, userRepo, req)
	response.ProcessMsgDto(c, msgDto)
}

func (h *Handler) Logout(c *gin.Context, req request.EmptyReq) {
	userOb := authentication.User{}
	jwt := &web.Jwt{GinContext: c}
	msgDto := userOb.Logout(jwt)
	response.ProcessMsgDto(c, msgDto)
}

func (h *Handler) Info(c *gin.Context, req request.EmptyReq) {
	userId := strconv.FormatUint(uint64(c.GetUint("userId")), 10)
	msgDto := dto.Msg{Success: true, Message: userId}
	response.ProcessMsgDto(c, msgDto)
}

func (h *Handler) Token(c *gin.Context, req request.EmptyReq) {
	userOb := authentication.User{}
	jwt := &web.Jwt{GinContext: c}
	msgDto := userOb.Token(jwt)
	response.ProcessMsgDto(c, msgDto)
}
