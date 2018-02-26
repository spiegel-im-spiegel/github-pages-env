package main

import (
	"log"
	"os"

	"github.com/hashicorp/logutils"
)

func main() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("WARN"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	log.Print("[DEBUG] Debugging")         // this will not print
	log.Print("[WARN] Warning")            // this will
	log.Print("[ERROR] Erring")            // and so will this
	log.Print("Message I haven't updated") // and so will this
}
