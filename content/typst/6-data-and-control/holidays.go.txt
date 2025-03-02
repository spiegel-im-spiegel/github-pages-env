package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/goark/koyomi"
	"github.com/goark/koyomi/value"
)

// Holiday represents a holiday with its date and title.
// Year is the year of the holiday.
// Month is the month of the holiday (1-12).
// Day is the day of the holiday (1-31).
// Weekday is the day of the week of the holiday.
// Title is the name or description of the holiday.
type Holiday struct {
	Year    int          `json:"year"`
	Month   int          `json:"month"`
	Day     int          `json:"day"`
	Weekday time.Weekday `json:"weekday"`
	Title   string       `json:"title"`
}

func main() {
	start, _ := value.DateFrom("2025-01-01")
	end, _ := value.DateFrom("2025-12-31")
	td, err := os.MkdirTemp(os.TempDir(), "blog")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer func() { _ = os.RemoveAll(td) }()
	k, err := koyomi.NewSource(
		koyomi.WithCalendarID(koyomi.Holiday),
		koyomi.WithStartDate(start),
		koyomi.WithEndDate(end),
		koyomi.WithTempDir(td),
	).Get()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	holidays := make([]Holiday, 0, len(k.Events()))
	for _, e := range k.Events() {
		holidays = append(holidays, Holiday{
			Year:    e.Date.Year(),
			Month:   int(e.Date.Month()),
			Day:     e.Date.Day(),
			Weekday: e.Date.Weekday(),
			Title:   e.Title,
		})
	}
	b, err := json.Marshal(holidays)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(string(b))
}
