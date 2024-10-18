package index

import (
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (eventh *Handler) Events(c *gin.Context) {
	//c.JSON(200, gin.H{"data": events})
}
