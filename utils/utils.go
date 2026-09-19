package utils

import (
	"fmt"
	"os"
	"path/filepath"

	appconfig "CustomerMS/config"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitUtils() error {
	configPath, err := configFilePath()
	if err != nil {
		return fmt.Errorf("locate configuration: %w", err)
	}
	if _, err := appconfig.Load(configPath); err != nil {
		return fmt.Errorf("validate configuration: %w", err)
	}

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read configuration: %w", err)
	}
	return nil
}

func EnvVarRead(key string) string {
	configPath, err := configFilePath()
	if err != nil {
		logrus.Errorf("Error locating configuration file for key %q: %v", key, err)
		return ""
	}

	values, err := godotenv.Read(configPath)
	if err != nil {
		logrus.Errorf("Error reading configuration file for key %q: %v", key, err)
		return ""
	}
	return values[key]
}

func configFilePath() (string, error) {
	if configuredPath := os.Getenv("BILLINGGO_CONFIG_FILE"); configuredPath != "" {
		if _, err := os.Stat(configuredPath); err != nil {
			return "", fmt.Errorf("configured path %q: %w", configuredPath, err)
		}
		return configuredPath, nil
	}

	workingDirectory, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get working directory: %w", err)
	}
	if path, err := findConfigFile(workingDirectory); err == nil {
		return path, nil
	}

	executable, err := os.Executable()
	if err == nil {
		if path, err := findConfigFile(filepath.Dir(executable)); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("configuration file config/configuration.env not found from working directory %q or executable directory", workingDirectory)
}

func findConfigFile(start string) (string, error) {
	current, err := filepath.Abs(start)
	if err != nil {
		return "", fmt.Errorf("resolve start directory: %w", err)
	}

	for {
		candidate := filepath.Join(current, "config", "configuration.env")
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}

		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}

	return "", fmt.Errorf("configuration file config/configuration.env not found from %q", start)
}
