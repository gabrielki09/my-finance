package config

import (
	"fmt"
	"os"
)

type ApiConfig struct {
	Port string
}

type DatabaseConfig struct {
	DbHost     string
	DbPort     string
	DbDatabase string
	DbUserName string
	DbPassword string
}

type Config struct {
	DatabaseConfig *DatabaseConfig
	ApiConfig      *ApiConfig
}

func getEnvValue(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("chave não localizada ou não definida: %s", key)
	}

	return value, nil
}

func loadApiConfig() (*ApiConfig, error) {
	port, err := getEnvValue("PORT")
	if err != nil {
		return nil, err
	}

	return &ApiConfig{
		Port: port,
	}, nil
}

func loadDbConfig() (*DatabaseConfig, error) {
	dbHost, err := getEnvValue("DB_HOST")
	if err != nil {
		return nil, err
	}

	dbPort, err := getEnvValue("DB_PORT")
	if err != nil {
		return nil, err
	}

	dbName, err := getEnvValue("DB_DATABASE")
	if err != nil {
		return nil, err
	}

	dbUser, err := getEnvValue("DB_USERNAME")
	if err != nil {
		return nil, err
	}

	dbPassword, err := getEnvValue("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	return &DatabaseConfig{
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbDatabase: dbName,
		DbUserName: dbUser,
		DbPassword: dbPassword,
	}, nil
}

func LoadConfig() (*Config, error) {
	dbConfig, err := loadDbConfig()
	if err != nil {
		return nil, err
	}

	apiConfig, err := loadApiConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		ApiConfig:      apiConfig,
		DatabaseConfig: dbConfig,
	}, nil
}
