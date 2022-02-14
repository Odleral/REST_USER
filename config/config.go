package config

import (
	"github.com/joeshaw/envdecode" //nolint:gci
	"log"
	"time" //nolint:gci
)

type Conf struct {
	Debug  bool `env:"DEBUG,required"`
	Server serverConf
	DB     dbConf
}

type serverConf struct {
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdel  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
	DBWait       time.Duration `env:"SERVER_WAIT_DB,required"`
}

type dbConf struct {
	Host     string `env:"DB_HOST_POSTGRES,required"`
	Port     int    `env:"DB_PORT_POSTGRES,required"`
	Username string `env:"DB_USER_POSTGRES,required"`
	Password string `env:"DB_PASSWORD_POSTGRES,required"`
	DBName   string `env:"DB_NAME_POSTGRES,required"`
}

func AppConfig() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("failed to decode: %s", err)
	}

	return &c
}
