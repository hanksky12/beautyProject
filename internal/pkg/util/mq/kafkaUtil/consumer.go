package kafkaUtil

import (
	"context"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Consumer struct {
	topic   string
	groupID string
	reader  *kafka.Reader
	wg      *sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
}

type MessageHandler interface {
	Handle(key string, value string, headers map[string]string)
}

func CreateConsumer(topic string, groupID string, wg *sync.WaitGroup) *Consumer {
	ctx, cancel := context.WithCancel(context.Background())
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{KafkaConf.Brokers[0].Address},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	//log.Infof("Consumer created [topic]: %s, [groupID]: %s\n", topic, groupID)
	return &Consumer{
		topic:   topic,
		groupID: groupID,
		reader:  reader,
		wg:      wg,
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (c *Consumer) Consume(handler MessageHandler) {
	defer c.wg.Done()

	for {
		select {
		case <-c.ctx.Done():
			log.Infof("Consumer shutting down [topic]: %s, [groupID]: %s", c.topic, c.groupID)
			return
		default:
			msg, err := c.reader.ReadMessage(c.ctx)
			if err != nil {
				if err == context.Canceled {
					return
				}
				log.Printf("Error reading message: %v\n", err)
				continue
			}
			headersMap := make(map[string]string)
			for _, header := range msg.Headers {
				headersMap[header.Key] = string(header.Value)
			}
			handler.Handle(string(msg.Key), string(msg.Value), headersMap)
			log.Infof("message at key=>%s msg=>%s\n", string(msg.Key), string(msg.Value))
		}
	}
}

func (c *Consumer) Stop() {
	c.cancel()
	c.reader.Close()
}

func ConsumeManualOffset(topic string, partition int, offset int64, callback func(string2 string, offset int64, key string, value string)) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{KafkaConf.Brokers[0].Address},
		Topic:     topic,
		Partition: partition,
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(offset) // 设置Offset

	// 接收消息
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		callback(topic, m.Offset, string(m.Key), string(m.Value))
		log.Infof("message at offset=>%d: key=>%s msg=>%s\n", m.Offset, string(m.Key), string(m.Value))
	}

	// 程序退出前关闭Reader
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
