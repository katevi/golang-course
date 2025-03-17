package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

const (
	logFile = "text.log"
)

func main() {
	file, err := os.Create(logFile)
	if err != nil {
		fmt.Println("failed to initialize file", logFile, err)
	}

	level := slog.LevelVar{}
	level.Set(slog.LevelDebug)

	handlerOptions := slog.HandlerOptions{
		AddSource: true,
		Level:     &level,
	}
	handler := slog.NewTextHandler(file, &handlerOptions)

	logger := slog.New(handler)

	logger.Debug("Debug message shown", slog.String("level_value", level.String()))
	logger.Info("Info message shown", slog.String("level_value", level.String()))
	logger.Warn("Warn message shown", slog.String("level_value", level.String()))
	logger.Error("Error message shown", slog.String("level_value", level.String()))

	level.Set(slog.LevelWarn)
	logger.Debug("Debug message not shown", slog.String("level_value", level.String()))
	logger.Info("Info message not shown", slog.String("level_value", level.String()))
	logger.Warn("Warn message shown", slog.String("level_value", level.String()))
	logger.Error("Error message shown", slog.String("level_value", level.String()))

	level.Set(slog.LevelDebug)
	logger.Debug("Stopping application", slog.Any("time", time.Now()))
}
