package router

import (
	"beautyProject/internal/pkg/util/web/gin/handler/response"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
)

func respRecover(c *gin.Context) {
	if err := recover(); err != nil {
		log.Infof("Recovered from panic: %v\n", err)
		log.Errorf("Stack trace:\n%s", debug.Stack())
		response.Panic(c)
	}
}
