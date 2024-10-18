package router

import (
	"beautyProject/internal/pkg/web/gin/handler/response"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

func respRecover(c *gin.Context) {
	if err := recover(); err != nil {
		log.Infof("Recovered from panic: %v\n", err)
		log.Criticalf("Stack trace:\n%s", debug.Stack())
		response.Panic(c)
	}
}
