package utils

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func EnvVarRead(key string) string {
	err := godotenv.Load("config/configuration.env")
	if err != nil {
		logrus.Errorln("Error Loading Configuration file")
		return ""
	}
	return os.Getenv(key)
}
