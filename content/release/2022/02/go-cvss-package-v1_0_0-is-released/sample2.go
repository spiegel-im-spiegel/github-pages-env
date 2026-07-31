//go:build run
// +build run

package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spiegel-im-spiegel/go-cvss/v3/metric"
	"github.com/spiegel-im-spiegel/go-cvss/v3/report"
	"golang.org/x/text/language"
)

var template = `- CVSS Version {{ .Version }}
- Vector: {{ .Vector }}

## Base Metrics

- Base Score: {{ .BaseScore }}

| {{ .BaseMetrics }} | {{ .BaseMetricValue }} |
|--------|-------|
| {{ .AVName }} | {{ .AVValue }} |
| {{ .ACName }} | {{ .ACValue }} |
| {{ .PRName }} | {{ .PRValue }} |
| {{ .UIName }} | {{ .UIValue }} |
| {{ .SName }} | {{ .SValue }} |
| {{ .CName }} | {{ .CValue }} |
| {{ .IName }} | {{ .IValue }} |
| {{ .AName }} | {{ .AValue }} |

## Temporal Metrics

- Temporal Score: {{ .TemporalScore }}
- {{ .SeverityName }}: {{ .SeverityValue }}

| {{ .TemporalMetrics }} | {{ .TemporalMetricValue }} |
|--------|-------|
| {{ .EName }} | {{ .EValue }} |
| {{ .RLName }} | {{ .RLValue }} |
| {{ .RCName }} | {{ .RCValue }} |

## Environmental Metrics

- {{ .SeverityName }}: {{ .SeverityValue }} ({{ .EnvironmentalScore }})

| {{ .EnvironmentalMetrics }} | {{ .EnvironmentalMetricValue }} |
|--------|-------|
| {{ .CRName }} | {{ .CRValue }} |
| {{ .IRName }} | {{ .IRValue }} |
| {{ .ARName }} | {{ .ARValue }} |
| {{ .MAVName }} | {{ .MAVValue }} |
| {{ .MACName }} | {{ .MACValue }} |
| {{ .MPRName }} | {{ .MPRValue }} |
| {{ .MUIName }} | {{ .MUIValue }} |
| {{ .MSName }}  | {{ .MSValue }} |
| {{ .MCName }}  | {{ .MCValue }} |
| {{ .MIName }}  | {{ .MIValue }} |
| {{ .MAName }}  | {{ .MAValue }} |
`

func main() {
	em, err := metric.NewEnvironmental().Decode("CVSS:3.1/AV:P/AC:H/PR:H/UI:N/S:U/C:H/I:H/A:H/E:F/RL:U/RC:C/CR:M/IR:H/AR:M/MAV:L/MAC:H/MPR:L/MUI:R/MS:U/MC:L/MI:H/MA:L") //Random CVSS Vector
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	r, err := report.NewEnvironmental(em, report.WithOptionsLanguage(language.Japanese)).ExportWith(strings.NewReader(template))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
