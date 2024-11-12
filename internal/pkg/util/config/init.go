package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func Init() {
	// 先從.env 讀取，沒有就從外部環境變數
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found, continuing...")
	}
	env := os.Getenv("GO_ENV")
	log.Printf(env)
	viper.SetConfigFile(fmt.Sprintf("internal/pkg/conf/config.%s.yaml", env))
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading default config file: %s", err)
	}
	fmt.Println("Config Init Success")
}
