package atomicbit

import (
	"testing"
)

func BenchmarkNewTrue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(true)
	}
}

func BenchmarkNewFalse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(false)
	}
}

func BenchmarkFlip(b *testing.B) {
	bit := New(true)
	for i := 0; i < b.N; i++ {
		bit.Flip()
	}
}

func BenchmarkGet(b *testing.B) {
	bit := New(true)
	for i := 0; i < b.N; i++ {
		bit.Get()
	}
}

func BenchmarkSetFalse(b *testing.B) {
	bit := New(true)
	for i := 0; i < b.N; i++ {
		bit.Set(false)
	}
}

func BenchmarkSetTrue(b *testing.B) {
	bit := New(true)
	for i := 0; i < b.N; i++ {
		bit.Set(true)
	}
}
