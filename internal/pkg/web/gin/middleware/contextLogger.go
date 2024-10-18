package middleware

import (
	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ContextLogger(router *gin.Engine) {
	router.Use(func(c *gin.Context) {
		requestID := uuid.New().String() // 生成請求 ID
		c.Set("requestID", requestID)    // 將請求 ID 存入上下文

		userID := c.GetHeader("X-User-ID") // 假設從 Header 獲取用戶 ID
		c.Set("userID", userID)            // 將用戶 ID 存入上下文

		userIP := c.ClientIP()  // 假設從 Header 獲取用戶 ID
		c.Set("userIP", userIP) // 將用戶 ID 存入上下文

		// 設置上下文日誌格式
		logFields := gin.H{
			"req_id":  requestID,
			"user_id": userID,
			"ip":      userIP,
		}
		// 記錄請求開始的日誌
		seelog.Info("start ", logFields)

		c.Next() // 讓請求繼續處理

		// 在請求處理完成後記錄日誌
		seelog.Info("end ", logFields)
	})
}
