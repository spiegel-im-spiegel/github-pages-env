package main

import (
	"fmt"
	"os"

	"github.com/goark/koyomi"
	"github.com/goark/koyomi/value"
)

func main() {
	start, _ := value.DateFrom("2026-01-01")
	end, _ := value.DateFrom("2026-12-31")
	td, err := os.MkdirTemp(os.TempDir(), "blog")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer func() { _ = os.RemoveAll(td) }()
	k, err := koyomi.NewSource(
		koyomi.WithCalendarID(
			koyomi.Holiday,
			// koyomi.MoonPhase,
			// koyomi.SolarTerm,
			// koyomi.Eclipse,
			// koyomi.Planet,
		),
		koyomi.WithStartDate(start),
		koyomi.WithEndDate(end),
		koyomi.WithTempDir(td),
	).Get()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("| 日付 | 曜日 | 内容 |")
	fmt.Println("| ---- |:----:| ---- |")
	for _, e := range k.Events() {
		fmt.Printf("| %v | %v | %v |\n", e.Date.StringJp(), e.Date.WeekdayJp().ShortStringJp(), e.Title)
	}
}

/* Copyright 2020-2025 Spiegel
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
