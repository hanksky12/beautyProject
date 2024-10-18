package response

import (
	"beautyProject/internal/pkg/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, message string, data any) {
	if message == "" {
		message = "請求成功"
	}
	base(c, dto.Response{Code: http.StatusOK, Message: message, Data: data})
}

func Failed(c *gin.Context, message string) {
	base(c, dto.Response{Code: http.StatusCreated, Message: message})
}

func Error(c *gin.Context, err error) {
	base(c, dto.Response{Code: http.StatusBadRequest, Message: err.Error()})
}

func Panic(c *gin.Context) {
	base(c, dto.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error"})
}

func ProcessMsgDto(c *gin.Context, msg dto.Msg) {
	if msg.Success == true {
		Success(c, "", gin.H{"message": msg.Message})
	} else {
		Failed(c, msg.Message)
	}
}

func base(c *gin.Context, response dto.Response) {
	c.JSON(http.StatusOK, response)
}
