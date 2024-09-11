package ilp

import (
	"runtime"
	"testing"
)

func BenchmarkIncrementInOrDependentValues(b *testing.B) {
	var a0, a1 int

	// go 1.21 443004858 ns/op 443 ms
	// go 1.23 481181190 ns/op 481 ms
	b.Run("increment independent values", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a0, a1 = IncrementIndependentValues()
		}
	})

	// go 1.21 100838541 ns 100 ms
	// go 1.23 444967265 ns 444 ms
	b.Run("increment dependent values", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a0, a1 = IncrementDependentValues()
		}
	})

	runtime.KeepAlive(a0)
	runtime.KeepAlive(a1)
}

func BenchmarkIncrement4InOrDependentValues(b *testing.B) {
	var a0, a1, a2, a3 int

	b.Run("increment independent values", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a0, a1, a2, a3 = Increment4IndependentValues()
		}
	})

	b.Run("increment dependent values", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a0, a1, a2, a3 = Increment4DependentValues()
		}
	})

	runtime.KeepAlive(a0)
	runtime.KeepAlive(a1)
	runtime.KeepAlive(a2)
	runtime.KeepAlive(a3)
}

func BenchmarkIncrement8InOrDependentValues(b *testing.B) {
	var a0, a1, a2, a3, a4, a5, a6, a7 int

	b.Run("increment independent values", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a0, a1, a2, a3, a4, a5, a6, a7 = Increment8IndependentValues()
		}
	})

	b.Run("increment dependent values", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a0, a1, a2, a3, a4, a5, a6, a7 = Increment8DependentValues()
		}
	})

	runtime.KeepAlive(a0)
	runtime.KeepAlive(a1)
	runtime.KeepAlive(a2)
	runtime.KeepAlive(a3)
	runtime.KeepAlive(a4)
	runtime.KeepAlive(a5)
	runtime.KeepAlive(a6)
	runtime.KeepAlive(a7)
}
