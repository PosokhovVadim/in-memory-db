package logger

import (
	"fmt"
	"in-memory-db/internal/config"
	"log/slog"
	"os"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func SetupLogger(cfg *config.Config) (*slog.Logger, error) {
	var log *slog.Logger

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		if err := os.Mkdir("logs", os.ModePerm); err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(fmt.Sprintf("logs/%s.log", "db"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	log = slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return log.With("env", cfg.Env), nil

}
