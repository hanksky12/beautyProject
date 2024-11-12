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
		userID := c.GetHeader("X-User-ID") // todo 從header中獲取用戶ID
		userIP := c.ClientIP()
		// 記錄請求開始的日誌
		requestLogger := log.WithFields(log.Fields{"user_id": userID, "user_ip": userIP})
		requestLogger.Info("[start]")
		c.Next()
		// 在請求處理完成後記錄日誌
		requestLogger.Info("[end]")
	})
}
