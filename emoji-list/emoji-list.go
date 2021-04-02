// +build run

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/spiegel-im-spiegel/emojis"
)

func dump(s string) string {
	ss := []string{}
	for _, r := range s {
		ss = append(ss, fmt.Sprintf("`%U`", r))
	}
	return strings.Join(ss, " ")
}

func main() {
	var flagShortcodes bool
	flag.BoolVar(&flagShortcodes, "s", false, "shortcodes")
	flag.Parse()

	list, err := emojis.NewEmojiSequenceList()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if flagShortcodes {
		fmt.Println("| Code Point | Shortcodes |")
		fmt.Println("| :--------: | ---------- |")

	} else {
		fmt.Println("| Code Point | Name | Type |")
		fmt.Println("| :--------: | ---- | ---- |")
	}
	for _, ec := range list {
		if flagShortcodes {
			if len(ec.Shortcodes) > 0 {
				var bldr strings.Builder
				for _, c := range ec.Shortcodes {
					bldr.WriteString(fmt.Sprintf(" `%s`", c))
				}
				fmt.Printf("| <span class='larger'><abbr class='emoji-chars' title='%v'>%v</abbr></span><br>%v |%v |\n", ec.Name, ec.Sequence, dump(ec.Sequence), bldr.String())
			}
		} else {
			fmt.Printf("| <span class='larger'><abbr class='emoji-chars' title='%[3]v'>%[1]v</abbr></span><br>%v | %v | `%v` |\n", ec.Sequence, dump(ec.Sequence), ec.Name, ec.SequenceType)
		}
	}
}
