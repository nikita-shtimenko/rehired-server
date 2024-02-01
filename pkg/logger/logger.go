package logger

import (
	"errors"
	"log/slog"
	"os"
)

var (
	ErrorEnvLogLevelNotFound = errors.New("Unable to match environment to the logging level. Invalid environment.")
)

var (
	logLevelMap = map[string]slog.Level{
		"dev":  slog.LevelDebug,
		"prod": slog.LevelInfo,
	}
)

func New(env string) (*slog.Logger, error) {
	_, ok := logLevelMap[env]

	if !ok {
		return nil, ErrorEnvLogLevelNotFound
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevelMap[env]}))
	return logger, nil
}
