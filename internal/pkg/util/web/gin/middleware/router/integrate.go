package router

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Integrate[T any](handlerFunc func(*gin.Context, T), auth Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer respRecover(c)
		ok, userID := authentication(c, auth)
		if !ok {
			return
		}
		userIP := c.ClientIP()
		log.Infof("user_id:%v user_ip:%v", userID, userIP)
		validator(handlerFunc)(c)
	}
}
