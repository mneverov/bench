package stack

import (
	"runtime"
	"testing"
)

func BenchmarkStacks(b *testing.B) {
	bStack := NewBlockingStack[int]()
	lfStack := NewLockFreeStack[int]()
	rStack := NewRegularStack[int]()
	var res int

	b.Run("sequential access regular stack", func(b *testing.B) {
		for i := range b.N {
			rStack.Push(i)
			res, _ = rStack.Pop()
		}
	})

	b.Run("sequential access lock free stack", func(b *testing.B) {
		for i := range b.N {
			lfStack.Push(i)
			res, _ = lfStack.Pop()
		}
	})

	b.Run("sequential access blocking stack", func(b *testing.B) {
		for i := range b.N {
			bStack.Push(i)
			res, _ = bStack.Pop()
		}
	})

	b.Run("parallel access lock free stack", func(b *testing.B) {
		i := 0
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				lfStack.Push(i)
				res, _ = lfStack.Pop()
				i++
			}
		})
	})

	b.Run("parallel access lock free stack writes only", func(b *testing.B) {
		i := 0
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				lfStack.Push(i)
				i++
			}
		})
	})

	b.Run("parallel access lock free stack reads only", func(b *testing.B) {
		i := 0
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				res, _ = lfStack.Pop()
				i++
			}
		})
	})

	b.Run("parallel access blocking stack", func(b *testing.B) {
		i := 0
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				bStack.Push(i)
				res, _ = bStack.Pop()
				i++
			}
		})
	})

	b.Run("parallel access blocking stack writes only", func(b *testing.B) {
		i := 0
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				bStack.Push(i)
				i++
			}
		})
	})

	b.Run("parallel access blocking stack reads only", func(b *testing.B) {
		i := 0
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				res, _ = bStack.Pop()
				i++
			}
		})
	})

	runtime.KeepAlive(res)
}

/*
goos: linux
goarch: amd64
pkg: github.com/mneverov/bench
cpu: 11th Gen Intel(R) Core(TM) i7-11850H @ 2.50GHz
BenchmarkStacks
BenchmarkStacks/sequential_access_regular_stack
BenchmarkStacks/sequential_access_regular_stack-16         	48885643	        27.34 ns/op                   1 goroutine no synchronization
BenchmarkStacks/sequential_access_lock_free_stack
BenchmarkStacks/sequential_access_lock_free_stack-16       	37588828	        32.78 ns/op                   1 goroutine CAS
BenchmarkStacks/sequential_access_blocking_stack
BenchmarkStacks/sequential_access_blocking_stack-16        	31150153	        38.69 ns/op                   1 goroutine mutex
BenchmarkStacks/parallel_access_lock_free_stack

BenchmarkStacks/parallel_access_lock_free_stack-16         	 6103536	       198.1 ns/op                    N goroutines CAS read/write combined
BenchmarkStacks/parallel_access_lock_free_stack_writes_only
BenchmarkStacks/parallel_access_lock_free_stack_writes_only-16         	10303651	       112.8 ns/op
BenchmarkStacks/parallel_access_lock_free_stack_reads_only
BenchmarkStacks/parallel_access_lock_free_stack_reads_only-16          	159716146	        16.92 ns/op
BenchmarkStacks/parallel_access_blocking_stack

BenchmarkStacks/parallel_access_blocking_stack-16                      	 8268328	       146.1 ns/op        N goroutines mutex read/write combined. Less than CAS
BenchmarkStacks/parallel_access_blocking_stack_writes_only
BenchmarkStacks/parallel_access_blocking_stack_writes_only-16          	17030862	        73.59 ns/op
BenchmarkStacks/parallel_access_blocking_stack_reads_only
BenchmarkStacks/parallel_access_blocking_stack_reads_only-16           	17163522	        73.84 ns/op
*/
