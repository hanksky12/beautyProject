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

func Unauthorized(c *gin.Context) {
	base(c, dto.Response{Code: http.StatusUnauthorized, Message: "Unauthorized"})
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

func ProcessTableDto(c *gin.Context, table dto.Table) {
	if table.Success == true {
		SuccessTable(c, table.DataList, table.Total)
	} else {
		FailedTable(c, table.Message)
	}
}

func SuccessTable(c *gin.Context, dataList []map[string]any, total int) {
	base(c, dto.Response{Code: http.StatusOK, Message: "請求成功", Data: dataList, Total: total})
}

func FailedTable(c *gin.Context, message string) {
	base(c, dto.Response{Code: http.StatusCreated, Message: message, Data: []map[string]any{}, Total: 0})
}

func base(c *gin.Context, response dto.Response) {
	c.JSON(http.StatusOK, response)
}
