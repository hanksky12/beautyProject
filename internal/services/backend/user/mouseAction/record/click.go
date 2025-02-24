package record

import (
	"beautyProject/internal/services/backend/common/record"
)

type Click struct {
	Recorder record.Recorder
}

func (c *Click) Work(userId string, msg []byte) {
	c.Recorder.Work(userId, msg)
}
