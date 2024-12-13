package main

import (
	"beautyProject/cmd/consumer/app"
	"beautyProject/internal/pkg/util/app/initApp"
	"beautyProject/internal/pkg/util/app/initApp/initBeforeApp"
)

func main() {
	options := &initApp.Options{
		KafkaOptions: &initApp.KafkaOptions{
			IsKafka:        true,
			IsCreateWriter: false,
		},
		LogOptions: &initApp.LogOptions{
			FilePath: "log/consumer.log",
		},
	}
	initBeforeApp.Run(options)
	app.Run()
}
