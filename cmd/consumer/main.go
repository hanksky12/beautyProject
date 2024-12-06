package main

import (
	"beautyProject/cmd/consumer/app"
	"beautyProject/internal/pkg/util/app/initApp/initAfterApp"
	"beautyProject/internal/pkg/util/app/initApp/initBeforeApp"
)

func main() {
	initBeforeApp.Run("log/consumer.log", false)
	app.Run()
	initAfterApp.Run(false)
}
