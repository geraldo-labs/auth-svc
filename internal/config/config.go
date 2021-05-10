package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

// Settings App configuration based on environments
type Settings struct {
	AppPort     int    `envconfig:"APP_PORT" default:"8080"`
	AppVersion  string `envconfig:"APP_VERSION" default:"dev"`
	DatabaseDSN string `envconfig:"DATABASE_DSN" default:"host=postgres user=root password=root dbname=pii port=5432 sslmode=disable TimeZone=Europe/Berlin"`
}

func New(prefix string, log *zap.Logger) *Settings {
	conf := &Settings{}
	if err := envconfig.Process(prefix, conf); err != nil {
		log.Fatal(fmt.Sprintf("%v", fmt.Errorf("%w", err)))
	}
	return conf
}
