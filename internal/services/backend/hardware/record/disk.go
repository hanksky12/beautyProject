package record

import (
	"beautyProject/internal/pkg/enum"
	"beautyProject/internal/pkg/util/mq/kafkaUtil"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/shirou/gopsutil/disk"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Disk struct {
}

func (d *Disk) Work(userId string) {
	log.Info("DiskWork is working...")
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	strPercent := strconv.FormatFloat(diskInfo.UsedPercent, 'f', -1, 64)
	message := fmt.Sprintf("DiskWork percent: %v", strPercent)
	log.Info(message)
	strTime := strconv.FormatInt(time.Now().Unix(), 10)
	ok := kafkaUtil.Produce(
		enum.Disk.Name,
		kafka.Message{
			//Key:   []byte(userId),
			Value: []byte(strPercent),
			Headers: []kafka.Header{
				{
					Key:   "user_id",
					Value: []byte(userId),
				},
				{
					Key:   "time",
					Value: []byte(strTime),
				},
			},
		})
	if !ok {
		log.Panic("DiskWork send message failed")
	}
}

func (d *Disk) Name() string {
	return "disk"
}
