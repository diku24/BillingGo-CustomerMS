package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitUtils() {
	viper.SetConfigFile("/config/configuration.env")
	if configVar := viper.ReadInConfig(); configVar != nil {
		logrus.Errorln("Error Reading Configurations Variables")
	}

}

func EnvVarRead(key string) string {
	var filePath string
	//err := godotenv.Load("config/configuration.env")
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		logrus.Errorf("Eroor Loading Current working directory: %v", err)
	}

	//filePath = fmt.Sprintf("%v\\config\\configuration.env", currentWorkingDirectory)
	filePath = fmt.Sprintf("%v/config/configuration.env", currentWorkingDirectory)

	err = godotenv.Load(filePath)
	if err != nil {
		logrus.Errorf("Error Loading Configuration file: %v For the Key: %v", err, key)
		return err.Error()
	}
	return os.Getenv(key)
}
