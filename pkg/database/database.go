package database

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New ...
func New(dsn string, log *zap.Logger) *gorm.DB {
	var db *gorm.DB
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Sprintf("%v", fmt.Errorf("not possible to connect to Postgres: %w", err)))
	}

	return db
}
