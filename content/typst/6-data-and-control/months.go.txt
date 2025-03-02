package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Month represents a calendar month with its associated year, month number,
// the first weekday of the month, and the last day of the month.
// Year is the year of the month.
// Month is the month number (1-12).
// FirstWeekday is the weekday of the first day of the month (0-6, where 0 is Sunday).
// Lastday is the last day of the month.
type Month struct {
	Year         int `json:"year"`
	Month        int `json:"month"`
	FirstWeekday int `json:"first_weekday"`
	Lastday      int `json:"lastday"`
}

func main() {
	year := 2025
	months := make([]Month, 0, 12)
	for month := 1; month <= 12; month++ {
		m := Month{
			Year:         year,
			Month:        month,
			FirstWeekday: int(time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC).Weekday()),
			Lastday:      time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day(),
		}
		months = append(months, m)
	}
	b, err := json.Marshal(months)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(string(b))
}
