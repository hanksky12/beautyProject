package main

import (
	"beautyProject/cmd/consumer/app"
	"beautyProject/internal/pkg/util/app/initAfterApp"
	"beautyProject/internal/pkg/util/app/initBeforeApp"
)

func main() {
	initBeforeApp.Run("log/router.log", false)
	app.Run()
	initAfterApp.Run(false)
}
