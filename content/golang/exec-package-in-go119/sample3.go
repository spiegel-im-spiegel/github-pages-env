package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	os.Setenv("NoDefaultCurrentDirectoryInExePath", "1")
	cmd := exec.Command("gpgpdump.exe", "version")
	if cmd.Err != nil {
		fmt.Println(cmd.Err)
		if !errors.Is(cmd.Err, exec.ErrDot) {
			return
		}
		cmd.Err = nil
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("output by %v:\n%v\n", cmd, string(out))
}
