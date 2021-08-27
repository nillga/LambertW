package lambertw

import (
	"math"
	"testing"
)

func BenchmarkVeberic(b *testing.B) {
	for a := 0.0; a < 3; a++ {
		b.Run("", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				W(-1,-math.Exp(-a))
			}
		})
		b.Run("", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				W(0,-math.Exp(-a))
			}
		})
		b.Run("", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				W(0,math.Exp(a))
			}
		})
	}
}
func BenchmarkFukushima(b *testing.B) {
	for a := 0.0; a < 3; a++ {
		b.Run("", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Fukushima(-1,-math.Exp(-a))
			}
		})
		b.Run("", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Fukushima(0,-math.Exp(-a))
			}
		})
		b.Run("", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Fukushima(0,math.Exp(a))
			}
		})
	}
}