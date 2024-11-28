package pc

import (
	"beautyProject/internal/pkg/util/mq/kafkaUtil"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/shirou/gopsutil/disk"
	log "github.com/sirupsen/logrus"
)

type Disk struct {
}

func (d *Disk) Work(userId string) {
	log.Info("DiskWork is working...")
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	message := fmt.Sprintf("DiskWork percent: %v", diskInfo.UsedPercent)
	ok := kafkaUtil.Produce(
		"disk",
		kafka.Message{
			//Key:   []byte(userId),
			Value: []byte(message),
			Headers: []kafka.Header{
				{
					Key:   "user_id",
					Value: []byte(userId),
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
