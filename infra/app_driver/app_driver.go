package app_driver

import (
	"github.com/joho/godotenv"
	"os"
)

type EnvAppDriver struct {
	AppPort string
}

func NewAppDriver() (EnvAppDriver, error) {
	if err := godotenv.Load(".env"); err != nil {
		return EnvAppDriver{}, err
	}

	appPort := os.Getenv("APP_PORT")

	return EnvAppDriver{
		AppPort: appPort,
	}, nil
}
