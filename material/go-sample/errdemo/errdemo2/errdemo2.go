package errdemo2

import (
	"os"

	"github.com/pkg/errors"
)

// F returns error.
func F() error {
	file, err := os.Open("not-exist.txt")
	if err != nil {
		return errors.Wrap(err, "Error by F() function")
	}
	defer file.Close()

	return nil
}
