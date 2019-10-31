package zerotime_test

import (
	"testing"
	"time"

	"zerotime"
)

func BenchmarkTrancateByFmt(b *testing.B) {
	loc, _ := time.LoadLocation("America/New_York")
	t := time.Now().In(loc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = zerotime.TrancateByFmt(t)
	}
}

func BenchmarkTrancateByFormat(b *testing.B) {
	loc, _ := time.LoadLocation("America/New_York")
	t := time.Now().In(loc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = zerotime.TrancateByFormat(t)
	}
}

func BenchmarkTrancateByCalc(b *testing.B) {
	loc, _ := time.LoadLocation("America/New_York")
	t := time.Now().In(loc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = zerotime.TrancateByCalc(t)
	}
}

func BenchmarkTrancateBydate(b *testing.B) {
	loc, _ := time.LoadLocation("America/New_York")
	t := time.Now().In(loc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = zerotime.TrancateByDate(t)
	}
}

func BenchmarkTrancateByUnixTime(b *testing.B) {
	loc, _ := time.LoadLocation("America/New_York")
	t := time.Now().In(loc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = zerotime.TrancateByUnixTime(t)
	}
}
