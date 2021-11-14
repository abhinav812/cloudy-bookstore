package config

import (
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

// TomlConfig - struct to wrap around log level and server configuration values read from config.toml
type TomlConfig struct {
	DB      database `toml:"database"`
	Server  serverInfo
	Logging loggingInfo
}

type database struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

type serverInfo struct {
	Port         int
	ReadTimeout  duration
	WriteTimeout duration
	IdleTimeout  duration
}

type loggingInfo struct {
	Debug bool
}

// AppTomlConfig - creates toml from environment variables
func AppTomlConfig() (*TomlConfig, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatalf("No config path specified")
	}

	var config TomlConfig

	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		log.Fatalf("Failed to load config.toml. %v", err)
		return nil, err
	}
	return &config, nil
}

type duration struct {
	time.Duration
}

// UnmarshalText - this is used in parsing duration values as string in toml files.
func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
