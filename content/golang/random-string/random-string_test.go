package randstr_test

import (
	"github.com/spiegel-im-spiegel/github-pages-env/content/golang/random-string"
	"testing"
	"time"
)

const (
	len64   = 64
	len128  = 128
	max512  = len64 * 8
	max1024 = len128 * 8
)

func TestCryptoString1(t *testing.T) {
	r := randstr.NewCryptoRandom()
	s := r.String(len64)
	if len(s) != len64 {
		t.Errorf("CryptoRandom.String(%v) = \"%v\"", len64, s)
	}
}

func BenchmarkRandomStringMath64t8(b *testing.B) {
	r := randstr.NewMathRandom(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < max512/len64; n++ {
			_ = r.String(len64)
		}
	}
}

func BenchmarkRandomStringMath512(b *testing.B) {
	r := randstr.NewMathRandom(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = r.String(max512)
	}
}

func BenchmarkRandomStringMath128t8(b *testing.B) {
	r := randstr.NewMathRandom(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < max1024/len128; n++ {
			_ = r.String(len128)
		}
	}
}

func BenchmarkRandomStringMath1024(b *testing.B) {
	r := randstr.NewMathRandom(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = r.String(max1024)
	}
}

func BenchmarkRandomStringCrypto64t8(b *testing.B) {
	r := randstr.NewCryptoRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < max512/len64; n++ {
			_ = r.String(len64)
		}
	}
}

func BenchmarkRandomStringCrypto512(b *testing.B) {
	r := randstr.NewCryptoRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = r.String(max512)
	}
}

func BenchmarkRandomStringCrypto128t8(b *testing.B) {
	r := randstr.NewCryptoRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < max1024/len128; n++ {
			_ = r.String(len128)
		}
	}
}

func BenchmarkRandomStringCrypto1024(b *testing.B) {
	r := randstr.NewCryptoRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = r.String(max1024)
	}
}
