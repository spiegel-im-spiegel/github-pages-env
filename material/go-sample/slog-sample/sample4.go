package main

import (
	"context"
	"os"

	"github.com/goark/errs"
	"golang.org/x/exp/slog"
)

func checkFileOpen(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errs.Wrap(
			err,
			errs.WithContext("path", path),
		)
	}
	defer file.Close()
	return nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr)).WithContext(context.TODO())
	logger.Enabled(slog.DebugLevel)

	if err := checkFileOpen("not-exist.txt"); err != nil {
		logger.Error("open error", err, slog.Any("info", err))
	} else {
		logger.Info("no error")
	}
}
