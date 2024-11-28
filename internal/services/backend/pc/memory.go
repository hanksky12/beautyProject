package pc

import (
	"beautyProject/internal/pkg/util/mq/kafkaUtil"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
)

type Memory struct {
}

func (m *Memory) Work(userId string) {
	log.Info("MemoryWork is working...")
	memInfo, _ := mem.VirtualMemory()
	message := fmt.Sprintf("MemoryWork percent: %v", memInfo.UsedPercent)
	ok := kafkaUtil.Produce(
		"memory",
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
		log.Panic("MemoryWork send message failed")
	}
}

func (m *Memory) Name() string {
	return "memory"
}
