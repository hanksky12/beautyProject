package router

import (
	"beautyProject/internal/pkg/util/web/gin/handler/response"
	"beautyProject/internal/pkg/web/request"
	"github.com/gin-gonic/gin"
	"reflect"
)

/*
Validator 是一個通用的驗證中間件，它會將綁定好的數據直接傳給處理器函數
*/
func validator[T any](handlerFunc func(*gin.Context, T)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqData T
		// 如果不是 request.EmptyReq 才進行綁定檢查
		if reflect.TypeOf(reqData) != reflect.TypeOf(request.EmptyReq{}) {
			if err := c.ShouldBindJSON(&reqData); err != nil {
				response.Error(c, err)
				c.Abort()
				return
			}
		}
		handlerFunc(c, reqData)
	}
}
