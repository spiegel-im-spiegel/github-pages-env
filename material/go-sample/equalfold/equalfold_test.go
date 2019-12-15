package equalfold

import (
	"strings"
	"testing"
)

var (
	lefts   = []string{"go", "ｇｏ"}
	rights  = []string{"Go", "GO", "go", "Ｇｏ", "ＧＯ", "ｇｏ"}
	rights2 = []string{"go", "go", "go", "ｇｏ", "ｇｏ", "ｇｏ"}
)

func BenchmarkEqualCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, left := range lefts {
			for _, right := range rights2 {
				if left == right {
					_ = left
				} else {
					_ = right
				}
			}
		}
	}
}

func BenchmarkEqualLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, left := range lefts {
			for _, right := range rights {
				if left == strings.ToLower(right) {
					_ = left
				} else {
					_ = right
				}
			}
		}
	}
}

func BenchmarkEqualFold(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, left := range lefts {
			for _, right := range rights {
				if strings.EqualFold(left, right) {
					_ = left
				} else {
					_ = right
				}
			}
		}
	}
}
