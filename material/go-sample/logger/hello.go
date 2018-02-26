package main

import (
	"os"

	"github.com/spiegel-im-spiegel/logf"
)

func main() {
	logf.SetOutput(os.Stdout)
	logf.SetMinLevel(logf.WARN)

	logf.Debug("Debugging")   // this will not print
	logf.Print("Information") // this will not print
	logf.Warn("Warning")      // this will
	logf.Error("Erroring")    // and so will this
}
