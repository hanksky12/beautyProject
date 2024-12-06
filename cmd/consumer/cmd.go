package main

import (
	"beautyProject/cmd/consumer/app/controller"
	"beautyProject/internal/pkg/util/app/cmd"
	"beautyProject/internal/pkg/util/app/initApp/initBeforeApp"
)

func main() {
	initBeforeApp.Run("log/cmd.log", false)
	cmd.Execute(&controller.Callback{})
}
