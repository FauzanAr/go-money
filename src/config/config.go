package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type MappedConfig struct {
	AppPort						string
	PostgreHost					string
	PostgreMaxConnection		int
	PostgreMaxIdleConnection	int
}

var config MappedConfig

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file");
	}

	postgreMaxConn, _ := strconv.Atoi(os.Getenv("POSTGRE_MAX_CONNECTION"))
	postgreMaxIdleConn, _ := strconv.Atoi(os.Getenv("POSTGRE_MAX_IDLE_CONNECTION"))

	config = MappedConfig{
		AppPort: os.Getenv("APP_PORT"),
		PostgreHost: os.Getenv("POSTGRE_HOST"),
		PostgreMaxConnection: postgreMaxConn,
		PostgreMaxIdleConnection: postgreMaxIdleConn,
	}
}

func Get() *MappedConfig {
	return &config
}