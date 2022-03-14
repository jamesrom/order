package compare

import (
	"testing"
)

func BenchmarkLessFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessFloat(float32(1.1), float32(2.2))
	}
}

func BenchmarkLessFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessFloat(float64(1.1), float64(2.2))
	}
}

func BenchmarkLessInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessSimple(1, 2)
	}
}

func BenchmarkLessSlowInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Less(1, 2)
	}
}

func BenchmarkLessFloat32Over(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessFloat(float32(3.1), float32(2.2))
	}
}

func BenchmarkLessFloat64Over(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessFloat(float64(3.1), float64(2.2))
	}
}

func BenchmarkLessIntOver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessSimple(3, 2)
	}
}

func BenchmarkLessSlowFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessFloat(float32(1.1), float32(2.2))
	}
}

func BenchmarkLessSlowFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Less(float64(1.1), float64(2.2))
	}
}

func BenchmarkLessSlowIntOver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Less(3, 2)
	}
}

func BenchmarkLessSlowFloat32Over(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LessFloat(float32(3.1), float32(2.2))
	}
}

func BenchmarkLessSlowFloat64Over(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Less(float64(3.1), float64(2.2))
	}
}
