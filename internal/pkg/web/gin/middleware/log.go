package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AddLog(router *gin.Engine) {
	gin.DefaultWriter = logWriter{}      // 輸出正常日誌
	gin.DefaultErrorWriter = logWriter{} // 輸出錯誤日誌

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logMessage := fmt.Sprintf("%s | %d | %s | %s | \"%s %s\"\n",
			param.ClientIP,
			param.StatusCode,
			param.Latency,
			param.Request.Method,
			param.Request.RequestURI,
			param.ErrorMessage,
		)
		return logMessage
	}))
}

type logWriter struct{}

// 實現 Write 方法，讓 Gin 使用 log 進行日誌記錄
func (l logWriter) Write(p []byte) (n int, err error) {
	logger := log.WithFields(log.Fields{"type": "gin"})
	logger.Info(string(p)) // 將日誌內容重定向到 log，這裡使用 info 級別
	return len(p), nil
}
