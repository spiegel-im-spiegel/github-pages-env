package errdemo1

import "os"

// F returns error.
func F() error {
	file, err := os.Open("not-exist.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
