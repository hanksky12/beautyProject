package initBeforeApp

import (
	"beautyProject/internal/pkg/util/config"
	"beautyProject/internal/pkg/util/db/sql"
	"beautyProject/internal/pkg/util/log"
	"beautyProject/internal/pkg/util/mq/kafkaUtil"
)

func Run(filePath string, createWriter bool) {
	config.Init()
	log.Init(filePath)
	sql.Init()
	kafkaUtil.Init()
	if createWriter {
		kafkaUtil.CreateProducer()
	}
}
