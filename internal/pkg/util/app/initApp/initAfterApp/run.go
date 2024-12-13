package initAfterApp

import (
	"beautyProject/internal/pkg/util/app/initApp"
	"beautyProject/internal/pkg/util/mq/kafkaUtil"
)

func Run(Options *initApp.Options) {
	if Options.KafkaOptions.IsKafka {
		kafkaUtil.Init()
		if Options.KafkaOptions.IsCreateWriter {
			kafkaUtil.CloseWriter()
		}
	}
}
