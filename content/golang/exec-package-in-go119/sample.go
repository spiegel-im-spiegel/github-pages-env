//go:build run
// +build run

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := "gpgpdump.exe"
	out, err := exec.Command(cmd, "version").CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("output by %v:\n%v\n", cmd, string(out))
}
