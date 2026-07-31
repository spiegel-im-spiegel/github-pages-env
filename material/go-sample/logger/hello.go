package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spiegel-im-spiegel/logf"
)

func main() {
	logf.SetMinLevel(logf.WARN)
	rl, err := rotatelogs.New("./log.%Y%m%d%H%M.txt")
	if err != nil {
		logf.Fatal(err)
		return
	}
	logf.SetOutput(rl)

	logf.Debug("Debugging")   // this will not print
	logf.Print("Information") // this will not print
	logf.Warn("Warning")      // this will
	logf.Error("Erroring")    // and so will this
}
