package router

import (
	"github.com/gin-gonic/gin"
)

func Integrate[T any](handlerFunc func(*gin.Context, string, T)) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer respRecover(c)
		reqId := c.MustGet("requestID").(string)
		validator(reqId, handlerFunc)(c)
	}
}
