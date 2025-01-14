package record

import (
	"beautyProject/internal/pkg/enum"
	"beautyProject/internal/pkg/util/mq/kafkaUtil"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Memory struct {
}

func (m *Memory) Work(userId string) {
	log.Info("MemoryWork is working...")
	memInfo, _ := mem.VirtualMemory()
	strPercent := strconv.FormatFloat(memInfo.UsedPercent, 'f', -1, 64)
	message := fmt.Sprintf("MemoryWork percent: %v", strPercent)
	log.Info(message)
	strTime := strconv.FormatInt(time.Now().Unix(), 10)
	ok := kafkaUtil.Produce(
		enum.Memory.Name,
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
		log.Panic("MemoryWork send message failed")
	}
}

func (m *Memory) Name() string {
	return "memory"
}
