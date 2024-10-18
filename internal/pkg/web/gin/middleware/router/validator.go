package router

import (
	"beautyProject/internal/pkg/web/gin/handler/response"
	"github.com/gin-gonic/gin"
)

/*
Validator 是一個通用的驗證中間件，它會將綁定好的數據直接傳給處理器函數
*/
func validator[T any](reqId string, handlerFunc func(*gin.Context, string, T)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqData T
		if err := c.ShouldBindJSON(&reqData); err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}
		handlerFunc(c, reqId, reqData)
	}
}
