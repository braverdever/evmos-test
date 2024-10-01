package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	EvmosUrl string
}

func NewConfig() *Config {
	return &Config{}
}

func (config *Config) LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	config.EvmosUrl = os.Getenv("EVMOS_URL")
	if config.EvmosUrl == "" {
		return fmt.Errorf("please check environment variables")
	}

	return nil
}
