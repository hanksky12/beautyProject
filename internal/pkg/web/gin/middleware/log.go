package middleware

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

func AddLog(router *gin.Engine) {
	gin.DefaultWriter = seeLogWriter{}      // 輸出正常日誌
	gin.DefaultErrorWriter = seeLogWriter{} // 輸出錯誤日誌

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

type seeLogWriter struct{}

// 實現 Write 方法，讓 Gin 使用 Seelog 進行日誌記錄
func (sw seeLogWriter) Write(p []byte) (n int, err error) {
	seelog.Info(string(p)) // 將日誌內容重定向到 Seelog，這裡使用 info 級別
	return len(p), nil
}
