package utils

import (
	"github.com/syedmrizwan/go-grpc/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	PRODUCTION  = "prod"
	DEVELOPMENT = "dev"
)

var logger *zap.Logger

func init() {
	//todo add logger config to customize the logger
	switch env.Env.BuildEnv {
	case DEVELOPMENT:
		logger, _ = zap.NewDevelopment()
		return
	case PRODUCTION:
		logger, _ = zap.NewProduction()
		return
	default:
		logger = zap.NewExample()
		return
	}
}

func GetLogger() *zap.Logger {
	return logger
}

func BuildLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.Level.SetLevel(zapcore.InfoLevel)

	var err error
	logger, err := config.Build()
	if err != nil {
		panic("Failed to setup logger")
	}

	return logger
}
