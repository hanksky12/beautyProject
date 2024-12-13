package router

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Integrate[T any](handlerFunc func(*gin.Context, T), isAuth bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer respRecover(c)
		ok, userID := auth(c, isAuth)
		if !ok {
			return
		}
		userIP := c.ClientIP()
		log.Infof("[start] user_id:%v user_ip:%v", userID, userIP)
		validator(handlerFunc)(c)
	}
}
