package lowercase

import (
	"strings"
	"testing"
)

func isDevNull1(name string) bool {
	if len(name) != 3 {
		return false
	}
	if name[0]|0x20 != 'n' {
		return false
	}
	if name[1]|0x20 != 'u' {
		return false
	}
	if name[2]|0x20 != 'l' {
		return false
	}
	return true
}

func isDevNull2(name string) bool {
	if len(name) != 3 {
		return false
	}
	if name[0] != 'n' && name[0] != 'N' {
		return false
	}
	if name[1] != 'u' && name[1] != 'U' {
		return false
	}
	if name[2] != 'l' && name[2] != 'L' {
		return false
	}
	return true
}

func isDevNull3(name string) bool {
	return strings.EqualFold(name, "nul")
}

var tests = []struct {
	in     string
	result bool
}{
	{"nul", true},
	{"Nul", true},
	{"nui", false},
	{"lun", false},
	{"nulllllllllllllll", false},
	{"nuuuuuuuul", false},
	{strings.Repeat("N", 3000), false},
}

func test(b *testing.B, f func(string) bool) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			if got := f(test.in); got != test.result {
				b.Fatalf("want %v but got %v for %v", test.result, got, test.in)
			}
		}
	}
}

func BenchmarkS1(b *testing.B) {
	test(b, isDevNull1)
}

func BenchmarkS2(b *testing.B) {
	test(b, isDevNull2)
}

func BenchmarkS3(b *testing.B) {
	test(b, isDevNull3)
}
