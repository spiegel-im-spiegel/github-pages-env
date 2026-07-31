package pointer

import (
	"fmt"
	"testing"
)

type S struct {
	a, b, c int64
	d, e, f string
	g, h, i float64
}

//go:noinline
func (s S) stack(s1 S) {}

//go:noinline
func (s *S) heap(s1 *S) {}

func (s S) ValueA() int64 { return s.a }

func byCopy() S {
	return S{
		a: 1, b: 1, c: 1,
		d: "foo", e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func byPointer() *S {
	return &S{
		a: 1, b: 1, c: 1,
		d: "foo", e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func BenchmarkMemoryStack(b *testing.B) {
	var s S
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = byCopy()
	}
	b.StopTimer()
	_ = fmt.Sprintf("%v", s.a)
}

func BenchmarkMemoryHeap(b *testing.B) {
	var s *S
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = byPointer()
	}
	b.StopTimer()
	_ = fmt.Sprintf("%v", s.a)
}

func BenchmarkMemoryStack2(b *testing.B) {
	var s S
	var s1 S

	s = byCopy()
	s1 = byCopy()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000000; i++ {
			s.stack(s1)
		}
	}
}

func BenchmarkMemoryHeap2(b *testing.B) {
	var s *S
	var s1 *S

	s = byPointer()
	s1 = byPointer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000000; i++ {
			s.heap(s1)
		}
	}
}

type IS interface {
	ValueA() int64
}

func byInterface() IS {
	return S{
		a: 1, b: 1, c: 1,
		d: "foo", e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func BenchmarkMemoryBox(b *testing.B) {
	var s IS
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = byInterface()
	}
	b.StopTimer()
	_ = fmt.Sprintf("%v", s.ValueA())
}
