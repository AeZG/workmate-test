package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"strings"
)

type HTTPServerConfig struct {
	Port string `mapstructure:"port" validate:"required"`
	Host string `mapstructure:"host" validate:"required"`
}

type Config struct {
	HTTPServer HTTPServerConfig `mapstructure:"http-server"`
}

func LoadConfig(configPath string) (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("WORKMATE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	if err := validateConfig(&cfg); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return &cfg, nil
}

func validateConfig(cfg *Config) error {
	validate := validator.New()
	return validate.Struct(cfg)
}
