package main

import (
	"beautyProject/cmd/backend/router"
	"beautyProject/internal/pkg/config"
	"beautyProject/internal/pkg/db/sql"
	logCustom "beautyProject/internal/pkg/log"
	"beautyProject/internal/pkg/model"
)

func main() {
	config.Init()
	logCustom.Init("log/data.log")
	sql.Init()
	sql.Db.AutoMigrate(&model.User{})
	router.Run()

	//requestLogger := log.WithFields(log.Fields{"request_id": "123", "user_ip": "241"})
	//requestLogger.Info("something happened on that request")
	//requestLogger = log.WithFields(log.Fields{"aaaa": "123"})
	//requestLogger.Warn("something not great happened")

}
