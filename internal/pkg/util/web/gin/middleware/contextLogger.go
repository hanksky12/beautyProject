package middleware

import (
	logCustom "beautyProject/internal/pkg/util/log"
	"beautyProject/internal/pkg/util/str"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ContextLogger(router *gin.Engine) {
	log.AddHook(&logCustom.UidHook{})
	router.Use(func(c *gin.Context) {
		str.SetUid()
		log.Info("start request")
		c.Next()
		log.Info("end request")
	})
}
