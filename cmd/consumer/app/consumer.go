package app

import (
	"beautyProject/cmd/consumer/app/controller"
	"beautyProject/internal/pkg/util/mq/kafkaUtil"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
)

func Run() {
	var wg sync.WaitGroup

	consumers := make([]*kafkaUtil.Consumer, 0, len(kafkaUtil.KafkaConf.Topics))
	consumers = start(consumers, wg)
	waitShuttingDown()
	stop(consumers)
	wg.Wait()
	log.Infoln("All consumers stopped.")

}

func start(consumers []*kafkaUtil.Consumer, wg sync.WaitGroup) []*kafkaUtil.Consumer {
	for _, isRaw := range []bool{true, false} {
		callback := controller.Callback{IsRaw: isRaw}
		for _, topic := range kafkaUtil.KafkaConf.Topics {
			handler := &ConsumeHandler{}
			handler.SetHandle(topic.Name, callback)
			groupID := fmt.Sprintf("raw_%v", isRaw)
			for i := 0; i < topic.Partition; i++ {
				consumer := kafkaUtil.CreateConsumer(topic.Name, groupID, &wg)
				log.Infof("Consumer created [topic]: %s, [groupID]: %s number %d", topic.Name, groupID, i)
				consumers = append(consumers, consumer)
				go consumer.Consume(handler)
				wg.Add(1)
			}
		}
	}
	return consumers
}

func stop(consumers []*kafkaUtil.Consumer) {
	for _, consumer := range consumers {
		consumer.Stop()
	}
}

func waitShuttingDown() {
	// 監聽關閉信號
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	<-sigChan // 等待信號
	log.Infoln("Shutting down consumers...")
}

//callback := controller.Callback{Minutes: 1}
//kafkaUtil.Consume("cpu", 0, 0, callback.CpuRecord)
