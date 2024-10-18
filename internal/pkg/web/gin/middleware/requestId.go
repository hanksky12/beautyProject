package middleware

import (
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

func RequestID(router *gin.Engine) {
	router.Use(func(c *gin.Context) {
		requestID := c.GetString("requestID")
		log.Infof("requestID123: %s", requestID)
		// 設置 seelog 的 logger，將 requestID 作為前綴打印
		logger, _ := log.CloneLogger(log.Current)
		logger.SetAdditionalStackDepth(0)
		//logger.SetContext(fmt.Sprintf("requestID: %s", "12312"))

		//logger.SetContext("requestID: " + requestID)
		//logger.SetContext(func() string { return fmt.Sprintf("requestID: %s", requestID) })

		// 使用自定義的 logger 來打印日誌
		log.ReplaceLogger(logger)

		// 繼續處理請求
		c.Next()

		// 請求結束後恢復原始 logger
		log.Flush()
	})
}
