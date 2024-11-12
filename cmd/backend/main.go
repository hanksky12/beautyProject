package main

import (
	"beautyProject/cmd/backend/app"
	"beautyProject/internal/pkg/util/app/initBeforeApp"
)

func main() {
	initBeforeApp.Run("log/router.log")
	//sql.Db.AutoMigrate(&model.User{}) //todo delete
	app.Run()
}

//requestLogger := log.WithFields(log.Fields{"request_id": "123", "user_ip": "241"})
//requestLogger.Info("something happened on that request")
//requestLogger = log.WithFields(log.Fields{"aaaa": "123"})
//requestLogger.Warn("something not great happened")
