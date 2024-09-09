package config

import (
	"asetku-bukan-asetmu/utils/common"
	"fmt"
	"os"
)

type DBConfig struct {
	Host, Port, Name, User, Password, Driver string
}

type APIConfig struct {
	APIHost, APIPort string
}

type Config struct {
	APIConfig
	DBConfig
	FileConfig
}

type FileConfig struct {
	FilePath string
}

func (c *Config) ReadConfig() error {
	err := common.LoadENV()
	if err != nil {
		return err
	}

	c.DBConfig = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.APIConfig = APIConfig{
		APIHost: os.Getenv("API_HOST"),
		APIPort: os.Getenv("API_PORT"),
	}

	c.FileConfig = FileConfig{
		FilePath: os.Getenv("FILE_PATH"),
	}

	if c.DBConfig.Host == "" || c.DBConfig.Port == "" || c.DBConfig.Name == "" || c.DBConfig.User == "" || c.DBConfig.Password == "" || c.DBConfig.Driver == "" || c.APIConfig.APIHost == "" || c.APIConfig.APIPort == "" {
		return fmt.Errorf("missing required enivronment variables")
	}

	return nil
}

func NewConfig() (*Config, error) {
	config := &Config{}
	err := config.ReadConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}
