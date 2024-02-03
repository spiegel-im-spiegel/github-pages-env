package main

import (
	"fmt"
	"os"

	"github.com/goark/koyomi"
	"github.com/goark/koyomi/value"
)

var weekShortNames = [7]string{"日", "月", "火", "水", "木", "金", "土"}

func WeekShortNameJp(dt value.DateJp) string {
	return weekShortNames[dt.Weekday()%7]
}

func main() {
	start, _ := value.DateFrom("2025-01-01")
	end, _ := value.DateFrom("2025-12-31")
	k, err := koyomi.NewSource(
		koyomi.WithCalendarID(
			// koyomi.Holiday,
			// koyomi.MoonPhase,
			koyomi.SolarTerm,
			// koyomi.Eclipse,
			// koyomi.Planet,
		),
		koyomi.WithStartDate(start),
		koyomi.WithEndDate(end),
	).Get()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("| 日付 | 曜日 | 内容 |")
	fmt.Println("| ---- |:----:| ---- |")
	for _, e := range k.Events() {
		fmt.Printf("| %v | %v | %v |\n", e.Date, WeekShortNameJp(e.Date), e.Title)
	}
}

/* Copyright 2020-2023 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
