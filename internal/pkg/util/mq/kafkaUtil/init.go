package kafkaUtil

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	env "github.com/spf13/viper"
	"net"
	"strconv"
)

var (
	WriterMap map[string]kafka.Writer
	KafkaConf KafkaConfig
)

func Init() {
	err := env.UnmarshalKey("Kafka", &KafkaConf)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	//log.Info(KafkaConf.Brokers[0].Address)
	createTopic()
	log.Info("Kafka Init Success")
}

func createTopic() {
	address := KafkaConf.Brokers[0].Address
	topics := KafkaConf.Topics
	err, controller := connectBroker(address)
	var controllerConn *kafka.Conn
	// 连接至leader节点
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	for _, topic := range topics {
		topicConfigs := []kafka.TopicConfig{
			{
				Topic:             topic.Name,
				NumPartitions:     topic.Partition,
				ReplicationFactor: topic.Replication,
			},
		}

		// 创建topic
		err = controllerConn.CreateTopics(topicConfigs...)
		if err != nil {
			panic(err.Error())
		}
	}
}

func connectBroker(address string) (error, kafka.Broker) {
	// 连接至任意kafka节点
	conn, err := kafka.Dial("tcp", address)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	// 获取当前控制节点信息
	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	return err, controller
}
