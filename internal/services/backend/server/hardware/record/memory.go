package record

import (
	"beautyProject/internal/services/backend/common/record"
	"github.com/shirou/gopsutil/mem"
	"strconv"
)

type Memory struct {
	Hardware record.Recorder
}

func (m *Memory) Work(userId string) {
	memInfo, _ := mem.VirtualMemory()
	strPercent := strconv.FormatFloat(memInfo.UsedPercent, 'f', -1, 64)
	bytePercent := []byte(strPercent)
	m.Hardware.Work(userId, bytePercent)
}

func (m *Memory) Name() string {
	return "memory"
}
