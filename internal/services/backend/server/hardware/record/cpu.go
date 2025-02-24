package record

import (
	"beautyProject/internal/services/backend/common/record"
	"github.com/shirou/gopsutil/cpu"
	"strconv"
	"time"
)

type Cpu struct {
	Hardware record.Recorder
}

func (c *Cpu) Work(userId string) {
	percent, _ := cpu.Percent(time.Second, false)
	strPercent := strconv.FormatFloat(percent[0], 'f', -1, 64)
	bytePercent := []byte(strPercent)
	c.Hardware.Work(userId, bytePercent)
}

func (c *Cpu) Name() string {
	return "cpu"
}
