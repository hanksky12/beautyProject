package main

import (
	"beautyProject/cmd/migrate/app/controller"
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
			FilePath: "log/migrate.log",
		},
	}
	initBeforeApp.Run(options)
	cmd.Execute(&controller.Command{})
	//go run ./cmd/migrate/cmd.go -c Run -p up -p stepsOrVersion
}
