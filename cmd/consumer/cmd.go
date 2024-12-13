package main

import (
	"beautyProject/cmd/consumer/app/controller"
	"beautyProject/internal/pkg/util/app/cmd"
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
			FilePath: "log/cmd.log",
		},
	}
	initBeforeApp.Run(options)
	cmd.Execute(&controller.Callback{})
}
