package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

type serverConf struct {
	Port         int           `env:"PORT, required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
}

// Conf - struct to wrap around log level and server configuration values
type Conf struct {
	Debug  bool `env:"DEBUG,required"`
	Server serverConf
}

// AppConfig - creates Conf from environment variables
func AppConfig() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}
	return &c
}
