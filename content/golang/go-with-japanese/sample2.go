// +build run

package main

import (
	"flag"
	"fmt"
	"unicode"
)

var mapRangeTable = map[*unicode.RangeTable]string{
	unicode.Variation_Selector: "Variation Selector",
	unicode.Sc:                 "Symbol/currency",
	unicode.Sk:                 "Symbol/modifier",
	unicode.Sm:                 "Symbol/math",
	unicode.Radical:            "Radical",
	unicode.So:                 "Symbol/other",
	unicode.Lm:                 "Letter/modifier",
	unicode.Ideographic:        "Ideographic",
	unicode.Lo:                 "Letter/other",
	unicode.Nl:                 "Number/letter",
	unicode.No:                 "Number/other",
	unicode.Mc:                 "Mark/spacing combining",
	unicode.Me:                 "Mark/enclosing",
	unicode.Mn:                 "Mark/nonspacing",
	unicode.Pc:                 "Punctuation/connector",
	unicode.Pd:                 "Punctuation/dash",
	unicode.Pe:                 "Punctuation/close",
	unicode.Pf:                 "Punctuation/final quote",
	unicode.Pi:                 "Punctuation/initial quote",
	unicode.Ps:                 "Punctuation/open",
	unicode.Po:                 "Punctuation/other",
	unicode.Zl:                 "Separator/line",
	unicode.Zp:                 "Separator/paragraph",
	unicode.Zs:                 "Separator/space",
	unicode.Join_Control:       "Join Control",
	unicode.Cc:                 "Control/control",
	unicode.Cf:                 "Control/format",
	unicode.Cs:                 "Control/surrogate",
	unicode.Co:                 "Control/private use",
}

var arryRangeTable = []*unicode.RangeTable{
	unicode.Variation_Selector,
	unicode.Sc,
	unicode.Sk,
	unicode.Sm,
	unicode.Radical,
	unicode.So,
	unicode.Lm,
	unicode.Ideographic,
	unicode.Lo,
	unicode.Nl,
	unicode.No,
	unicode.Mc,
	unicode.Me,
	unicode.Mn,
	unicode.Pc,
	unicode.Pd,
	unicode.Pe,
	unicode.Pf,
	unicode.Pi,
	unicode.Ps,
	unicode.Po,
	unicode.Zl,
	unicode.Zp,
	unicode.Zs,
	unicode.Join_Control,
	unicode.Cc,
	unicode.Cf,
	unicode.Cs,
	unicode.Co,
}

type unicodeChecker struct {
	mapRangeTable  map[*unicode.RangeTable]string
	arryRangeTable []*unicode.RangeTable
}

func newChecker() *unicodeChecker {
	return &unicodeChecker{mapRangeTable: mapRangeTable, arryRangeTable: arryRangeTable}
}

func (c *unicodeChecker) string(rt *unicode.RangeTable) string {
	if c == nil || rt == nil {
		return "Unknown"
	}
	if s, ok := c.mapRangeTable[rt]; ok {
		return s
	}
	return "Unknown"
}

func (c *unicodeChecker) check(r rune) string {
	// if unicode.IsGraphic(r) {
	// 	return "Graphic"
	// }
	if c == nil {
		return c.string(nil)
	}
	for _, rt := range c.arryRangeTable {
		if unicode.Is(rt, r) {
			return c.string(rt)
		}
	}
	return c.string(nil)
}

func main() {
	flag.Parse()
	args := flag.Args()
	checker := newChecker()
	for _, arg := range args {
		fmt.Println(arg)
		for i, r := range arg {
			switch {
			case unicode.IsLetter(r), r == '_':
				fmt.Printf("\t%#U : OK : letter\n", r)
			case unicode.IsNumber(r):
				okng := "OK"
				if i == 0 {
					okng = "NG"
				}
				fmt.Printf("\t%#U : %v : unicode_digit\n", r, okng)
			default:
				fmt.Printf("\t%#U : NG : %v\n", r, checker.check(r))
			}
		}
	}
}
