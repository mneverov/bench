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
