package main

import (
	"beautyProject/cmd/cron/app"
	"beautyProject/internal/pkg/util/app/initApp"
	"beautyProject/internal/pkg/util/app/initApp/initBeforeApp"
)

func main() {
	options := &initApp.Options{
		KafkaOptions: &initApp.KafkaOptions{
			IsKafka:        false,
			IsCreateWriter: false,
		},
		LogOptions: &initApp.LogOptions{
			FilePath: "log/cron.log",
		},
	}
	initBeforeApp.Run(options)
	app.Run()
}
