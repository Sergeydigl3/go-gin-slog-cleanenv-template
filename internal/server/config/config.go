package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

// Config represents the application's configuration.

type Config struct {
	Env string `yaml:"env" env-default:"local"`

	Server struct {
		Port int `yaml:"port" env-default:"9090"`
	} `yaml:"server"`

	//JWT struct {
	//	AccessSecret    string        `yaml:"secret" env-default:"secret"`
	//	RefreshSecret   string        `yaml:"refresh_secret" env-default:"refreshSecret"`
	//	TokenTTLAccess  time.Duration `yaml:"token_ttl" env-default:"1800s"`
	//	TokenTTLRefresh time.Duration `yaml:"token_ttl" env-default:"604800s"`
	//}
	//
	//MongoDB struct {
	//	ConnectUrl string `yaml:"connect_url" env-required:"true"`
	//	Database   string `yaml:"database" env-required:"true"`
	//	Collection string `yaml:"videos_collection" env-required:"true"`
	//} `yaml:"database"`
}

// MustLoad loads the configuration from the specified path and environment variables.
func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file doesn't exist: " + configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic(fmt.Errorf("failed to read config: %w", err))
	}

	return &cfg
}

// fetchConfigPath fetches the config path from command-line flags or environment variables.
// Priority: command-line flag > environment variable.
func fetchConfigPath() string {
	var configPath string
	flag.StringVar(&configPath, "config", "", "Path to the config file")
	flag.Parse()

	if configPath == "" {
		configPath = os.Getenv("CONFIG_PATH")
	}

	if configPath == "" {
		panic("no config file path provided")
	}

	return configPath
}
