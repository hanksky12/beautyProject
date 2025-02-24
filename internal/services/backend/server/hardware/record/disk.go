package record

import (
	"beautyProject/internal/services/backend/common/record"
	"github.com/shirou/gopsutil/disk"
	"strconv"
)

type Disk struct {
	Hardware record.Recorder
}

func (d *Disk) Work(userId string) {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	strPercent := strconv.FormatFloat(diskInfo.UsedPercent, 'f', -1, 64)
	bytePercent := []byte(strPercent)
	d.Hardware.Work(userId, bytePercent)
}

func (d *Disk) Name() string {
	return "disk"
}
