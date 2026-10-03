package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort       string
	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
	DatabaseName     string
	NetworkProtocol  string
}

func Load(path string) (Config, error) {
	values, err := godotenv.Read(path)
	if err != nil {
		return Config{}, fmt.Errorf("read configuration %q: %w", path, err)
	}

	value := func(key string) (string, error) {
		result := values[key]
		if result == "" {
			return "", fmt.Errorf("missing required configuration %q", key)
		}
		return result, nil
	}

	serverPort, err := value("SERVERPORT")
	if err != nil {
		return Config{}, err
	}
	databaseUser, err := value("DATABASEUSER")
	if err != nil {
		return Config{}, err
	}
	databasePassword, err := value("DATABASEPASS")
	if err != nil {
		return Config{}, err
	}
	databaseHost, err := value("DBCONNECTION")
	if err != nil {
		return Config{}, err
	}
	databaseName, err := value("DATABASE")
	if err != nil {
		return Config{}, err
	}
	networkProtocol, err := value("NETPROTOCOL")
	if err != nil {
		return Config{}, err
	}

	return Config{
		ServerPort:       serverPort,
		DatabaseUser:     databaseUser,
		DatabasePassword: databasePassword,
		DatabaseHost:     databaseHost,
		DatabaseName:     databaseName,
		NetworkProtocol:  networkProtocol,
	}, nil
}
