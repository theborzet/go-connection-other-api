package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string
	ApiKey     string
	ApiSecret  string
	RecvWindow string
}

func LOadConfig() (*Config, error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	port := viper.GetString("PORT")
	apiKey := viper.GetString("BYBIT_API_KEY")
	apiSecret := viper.GetString("BYBIT_API_SECRET")

	if apiKey == "" || apiSecret == "" {
		return nil, fmt.Errorf("API ключи не установлены в конфигурации")
	}

	config := &Config{
		Port:       port,
		ApiKey:     apiKey,
		ApiSecret:  apiSecret,
		RecvWindow: "10000",
	}

	return config, nil

}
