package main

import (
	"beautyProject/cmd/backend/app"
	"beautyProject/internal/pkg/util/app/initApp"
	"beautyProject/internal/pkg/util/app/initApp/initAfterApp"
	"beautyProject/internal/pkg/util/app/initApp/initBeforeApp"
)

func main() {
	options := &initApp.Options{
		KafkaOptions: &initApp.KafkaOptions{
			IsKafka:        true,
			IsCreateWriter: true,
		},
		LogOptions: &initApp.LogOptions{
			FilePath: "log/router.log",
		},
	}
	initBeforeApp.Run(options)
	app.Run()
	initAfterApp.Run(options)
}
