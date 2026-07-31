package main

import (
	"fmt"
	"strings"
)

type CharEncoding int

const (
	Unknown CharEncoding = iota
	UTF8
	ShiftJIS
	EUCJP
	ISO2022JP
)

var encodingMap = map[CharEncoding]string{
	UTF8:      "UTF-8",
	ShiftJIS:  "Shift_JIS",
	EUCJP:     "EUC-JP",
	ISO2022JP: "ISO-2022-JP",
}

func GetEncoding(s string) CharEncoding {
	for key, value := range encodingMap {
		if strings.EqualFold(value, s) {
			return key
		}
	}
	return Unknown
}

func (c CharEncoding) String() string {
	if s, ok := encodingMap[c]; ok {
		return s
	}
	return "Unknown"
}

func main() {
	fmt.Println(GetEncoding("utf-8"))
	fmt.Println(UTF8)
	//Output:
	//UTF-8
}
