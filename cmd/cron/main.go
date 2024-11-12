package main

import (
	"beautyProject/cmd/cron/app"
	"beautyProject/internal/pkg/util/app/initBeforeApp"
)

func main() {
	initBeforeApp.Run("log/cron.log")
	app.Run()
}
