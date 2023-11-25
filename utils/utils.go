package utils

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var filePath string = "C:/Users/Dinesh/go/src/BillingGo/CustomerMS/config/configuration.env"

func EnvVarRead(key string) string {
	//err := godotenv.Load("config/configuration.env")
	err := godotenv.Load(filePath)
	if err != nil {
		logrus.Errorf("Error Loading Configuration file: %v For the Key: %v", err, key)
		return err.Error()
	}
	return os.Getenv(key)
}
