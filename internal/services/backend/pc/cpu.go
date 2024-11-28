package pc

import (
	"beautyProject/internal/pkg/util/mq/kafkaUtil"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/shirou/gopsutil/cpu"
	log "github.com/sirupsen/logrus"
	"time"
)

type Cpu struct {
}

func (c *Cpu) Work(userId string) {
	log.Info("CpuWork is working...")
	percent, _ := cpu.Percent(time.Second, false)
	message := fmt.Sprintf("CpuWork percent: %v", percent[0])
	ok := kafkaUtil.Produce(
		"cpu",
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
		log.Panic("CpuWork send message failed")
	}
}

func (c *Cpu) Name() string {
	return "cpu"
}
