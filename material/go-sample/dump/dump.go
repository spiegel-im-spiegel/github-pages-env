package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("enc1")
	if err != nil {
		return
	}
	defer file.Close()

	b := make([]byte, 1)
	sep := "var byteList = []byte{"
	for true {
		if _, err = file.Read(b); err != nil {
			break
		}
		fmt.Printf("%s%#02x", sep, b)
		sep = ", "
	}
	fmt.Println("}")
}
