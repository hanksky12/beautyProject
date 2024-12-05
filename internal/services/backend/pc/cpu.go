package pc

import (
	"beautyProject/internal/pkg/enum"
	"beautyProject/internal/pkg/util/mq/kafkaUtil"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/shirou/gopsutil/cpu"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Cpu struct {
}

func (c *Cpu) Work(userId string) {
	log.Info("CpuWork is working...")
	percent, _ := cpu.Percent(time.Second, false)
	strPercent := strconv.FormatFloat(percent[0], 'f', -1, 64)
	message := fmt.Sprintf("CpuWork percent: %v", strPercent)
	log.Info(message)
	strTime := strconv.FormatInt(time.Now().Unix(), 10)
	ok := kafkaUtil.Produce(
		enum.Cpu.Name,
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
		log.Panic("CpuWork send message failed")
	}
}

func (c *Cpu) Name() string {
	return "cpu"
}
