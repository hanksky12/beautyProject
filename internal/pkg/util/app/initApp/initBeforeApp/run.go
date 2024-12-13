package initBeforeApp

import (
	"beautyProject/internal/pkg/util/app/initApp"
	"beautyProject/internal/pkg/util/config"
	"beautyProject/internal/pkg/util/db/sql"
	logCustom "beautyProject/internal/pkg/util/log"
	"beautyProject/internal/pkg/util/mq/kafkaUtil"
	log "github.com/sirupsen/logrus"
)

func Run(Options *initApp.Options) {
	config.Init()
	logCustom.Init(Options.LogOptions.FilePath)
	log.AddHook(&logCustom.GoroutineIDHook{})
	sql.Init()
	if Options.KafkaOptions.IsKafka {
		kafkaUtil.Init()
		if Options.KafkaOptions.IsCreateWriter {
			kafkaUtil.CreateProducer()
		}
	}
}

//func Run(filePath string, createWriter bool) {
//	config.Init()
//	logCustom.Init(filePath)
//	log.AddHook(&logCustom.GoroutineIDHook{})
//	sql.Init()
//	kafkaUtil.Init()
//	if createWriter {
//		kafkaUtil.CreateProducer()
//	}
//}
