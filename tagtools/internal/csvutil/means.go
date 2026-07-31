package csvutil

import (
	"errors"
	"io"
	"os"

	"github.com/goark/csvdata"
	"github.com/goark/errs"
)

// LoadMeansMap loads existing means values keyed by tag from tagslist CSV.
func LoadMeansMap(path string) (map[string]string, error) {
	meansMap := map[string]string{}

	f, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return meansMap, nil
		}
		return nil, errs.Wrap(err, errs.WithContext("path", path))
	}
	defer f.Close()

	rows := csvdata.NewRows(csvdata.New(f), true)
	defer rows.Close()

	for {
		err := rows.Next()
		if err != nil {
			if errors.Is(err, io.EOF) || errs.Is(err, io.EOF) {
				break
			}
			return nil, errs.Wrap(err, errs.WithContext("path", path))
		}

		tag := rows.Column("tag")
		if tag == "" {
			continue
		}
		meansMap[tag] = rows.Column("means")
	}

	return meansMap, nil
}
