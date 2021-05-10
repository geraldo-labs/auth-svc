package logging

import (
	"fmt"
	"go.uber.org/zap"
	"log"
)

func New() *zap.Logger {
	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatalf("error creating: %v", fmt.Errorf("%w", err))
	}

	defer logger.Sync()
	return logger
}
