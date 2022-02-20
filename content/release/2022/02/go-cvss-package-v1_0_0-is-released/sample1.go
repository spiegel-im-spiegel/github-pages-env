//go:build run
// +build run

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/spiegel-im-spiegel/go-cvss/v3/metric"
	"github.com/spiegel-im-spiegel/go-cvss/v3/report"
	"golang.org/x/text/language"
)

var template = `| {{ .BaseMetrics }} | {{ .BaseMetricValue }} |
|--------|-------|
| {{ .AVName }} | {{ .AVValue }} |
| {{ .ACName }} | {{ .ACValue }} |
| {{ .PRName }} | {{ .PRValue }} |
| {{ .UIName }} | {{ .UIValue }} |
| {{ .SName }} | {{ .SValue }} |
| {{ .CName }} | {{ .CValue }} |
| {{ .IName }} | {{ .IValue }} |
| {{ .AName }} | {{ .AValue }} |
`

func main() {
	bm, err := metric.NewBase().Decode("CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H") //CVE-2021-44716: net/http: limit growth of header canonicalization cache
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	// r, err := report.NewBase(bm).ExportWithString(template)
	r, err := report.NewBase(bm, report.WithOptionsLanguage(language.Japanese)).ExportWithString(template)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
