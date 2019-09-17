package statik

import (
	"net/http"

	"github.com/rakyll/statik/fs"
)

func New() (http.FileSystem, error) {
	return fs.New()
}
