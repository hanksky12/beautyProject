package main

import (
	"beautyProject/cmd/cron/app/controller"
	"beautyProject/internal/pkg/util/app/cmd"
	"beautyProject/internal/pkg/util/app/initBeforeApp"
)

func main() {
	initBeforeApp.Run("log/cmd.log")
	cmd.Execute(&controller.Job{})
}