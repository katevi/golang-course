package main

import (
	"log/slog"
	"os"

	"example.tester/coordinate"
)

func main() {
	input1 := "(42, 48, 11)"
	input2 := "((43, ,,)"

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	logger.Debug("Coordinate defined", slog.String("coordinate", input1))

	parsed1, err := coordinate.Parse(input1)
	if err != nil {
		logger.Error("Failed to parse coordinate", slog.String("coordinate", input1))
	} else {
		logger.Debug("Parsed coordinate successfully", slog.Any("parsed", parsed1))
	}

	parsed2, err := coordinate.Parse(input2)
	if err != nil {
		logger.Error("Failed to parse coordinate", slog.String("coordinate", input2))
	} else {
		logger.Debug("Parsed coordinate successfully", slog.Any("parsed", parsed2))
	}
}
