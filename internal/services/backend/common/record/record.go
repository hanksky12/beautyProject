package record

import (
	"beautyProject/internal/pkg/util/mq/kafkaUtil"
	"fmt"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Recorder struct {
	Name string
}

func (r *Recorder) Work(userId string, msg []byte) {
	log.Infof("%s Work is working...", r.Name)
	message := fmt.Sprintf("%s Work msg: %v", r.Name, msg)
	log.Info(message)
	strTime := strconv.FormatInt(time.Now().Unix(), 10)
	topic := r.Name
	ok := kafkaUtil.Produce(
		topic,
		kafka.Message{
			Value: msg,
			Headers: []kafka.Header{
				{
					Key:   "topic",
					Value: []byte(topic),
				},
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
		log.Panic(fmt.Sprintf("%s Work send message failed", r.Name))
	}
}
