package kafkaUtil

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"time"
)

func CreateProducer() {
	topics := KafkaConf.Topics
	address := KafkaConf.Brokers[0].Address
	WriterMap = make(map[string]kafka.Writer)
	for _, topic := range topics {
		//log.Info("topic: ", topic.Name)
		writer := kafka.Writer{
			Addr:         kafka.TCP(address),
			Topic:        topic.Name,
			Balancer:     &kafka.Hash{},    // 選擇分區的策略
			RequiredAcks: kafka.RequireAll, // 確保消息被所有同步副本確認
			//AllowAutoTopicCreation: true,             //topic不存在時自動創建
		}
		WriterMap[topic.Name] = writer
	}
}

func Produce(topic string, msg kafka.Message) bool {
	writer, ok := WriterMap[topic]
	if !ok {
		log.Errorf("failed to get writer: %v", topic)
		return false
	}

	var err error
	const retries = 3
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err = writer.WriteMessages(ctx, msg)
		cancel()
		if err == nil {
			//log.Printf("Message sent successfully: %s\n", msg.Value)
			return true // 成功後立即返回
		}
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			// 建立new topic 時，需要選舉
			time.Sleep(time.Millisecond * 250)
			continue
		}
		if err != nil {
			log.Fatalf("unexpected error %v", err)
			return false
		}
	}

	log.Printf("Message sent: %s\n", msg.Value)
	return true
}
