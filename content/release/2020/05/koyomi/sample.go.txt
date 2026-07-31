package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/spiegel-im-spiegel/koyomi"
)

func main() {
	start, _ := koyomi.DateFrom("2020-05-01")
	end, _ := koyomi.DateFrom("2020-05-31")
	k, err := koyomi.NewSource(
		koyomi.WithCalendarID(koyomi.Holiday, koyomi.SolarTerm),
		koyomi.WithStartDate(start),
		koyomi.WithEndDate(end),
	).Get()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	//b, err := k.EncodeCSV()
	b, err := k.EncodeJSON()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	io.Copy(os.Stdout, bytes.NewReader(b))
}
