package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/spiegel-im-spiegel/errs"
)

func checkFileOpen(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errs.Wrap(
			err,
			"file open error",
			errs.WithContext("path", path),
		)
	}
	defer file.Close()

	return nil
}

func main() {
	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().
		Timestamp().
		Str("role", "logger-sample").
		Logger()

	if err := checkFileOpen("not-exist.txt"); err != nil {
		//logger.Error().Err(err).Send()
		logger.Error().Interface("error", err).Send()
		//logger.Error().Interface("error", err).Send()
		//fmt.Printf("%+v\n", err)
	}
}
