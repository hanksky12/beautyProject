package main

import (
	"beautyProject/cmd/cron/app"
	"beautyProject/internal/pkg/util/app/initApp/initBeforeApp"
)

func main() {
	initBeforeApp.Run("log/cron.log", false)
	app.Run()
}
