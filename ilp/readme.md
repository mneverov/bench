# Instruction Level Parallelism

With [ILP](https://en.wikipedia.org/wiki/Instruction-level_parallelism) processors should be able to run independent
instructions simultaneously.

Function `IncrementDependentValues` increments the same element of an array `a[0]` `numIter` times, hence the
expectation is it should be slower than `IncrementIndependentValues` that does the same, but with independent elements.
Benchmark shows that it is not true:

```sh
go test -test.bench BenchmarkIncrementInOrDependentValues -test.trace=trace.out ./ilp/...
goos: linux
goarch: amd64
pkg: github.com/mneverov/bench/ilp
cpu: 11th Gen Intel(R) Core(TM) i7-11850H @ 2.50GHz
BenchmarkIncrementInOrDependentValues/increment_independent_values-16                  3         443918355 ns/op
BenchmarkIncrementInOrDependentValues/increment_dependent_values-16                    3         420649295 ns/op
PASS
ok      github.com/mneverov/bench/ilp   5.156s
```

i.e. incrementing depending values is faster.

If we compile the code to produce assembly with `go tool compile -S ilp/ilp.go > ilp.s` or use https://godbolt.org/ we
see that `IncrementDependentValues` is optimized to just add 2 instead of sequential incrementing:

```assembly
   ADDQ    $2, main.a(SP)     # a[0]+2
```

whereas `IncrementIndependentValues` does two increments:

```assembly
	INCQ    main.a(SP)       # a[0]++
	INCQ    main.a+8(SP)     # a[1]++
```

Let's check assembly code compiled with disabled optimizations (`go tool compile -S -l -N ilp/ilp.go > ilp.s`):

```assembly
	MOVQ    main.a(SP), CX # move value of a[0] to register CX 
	LEAQ    1(CX), DX      # load a[0]+1 to DX 🤷 
	MOVQ    DX, main.a(SP) # move DX which is a[0]+1 to a[0]
	ADDQ    $2, CX         # CX=a[0]+2
	MOVQ    CX, main.a(SP) # move value from CX back to a[0] 
```

Looks good (except 🤷 part). Run tests again with the disabled optimizations:

```sh
go test -gcflags='-N -l' -test.bench BenchmarkIncrementInOrDependentValues -test.trace=trace.out ./ilp/...
goos: linux
goarch: amd64
pkg: github.com/mneverov/bench/ilp
cpu: 11th Gen Intel(R) Core(TM) i7-11850H @ 2.50GHz
BenchmarkIncrementInOrDependentValues/increment_independent_values-16                  3         418164019 ns/op
BenchmarkIncrementInOrDependentValues/increment_dependent_values-16                    3         388270575 ns/op
PASS
ok      github.com/mneverov/bench/ilp   4.848s
```

Deoptimized version still runs faster. To me second and third assembly instructions (`LEAQ`, `MOVQ`) that increments
`a[0]` and stores it back do not make sense, because the next two `ADDQ` and `MOVQ` overwrite the value.

## Go v1.21

Compilation with go v1.21.13 produces the same code for `IncrementIndependentValues`, and also generates deoptimized
code for `IncrementDependentValues` -- same as for go `v1.23`.

Benchmark results:

```sh
$ go1.21.13 test -test.bench BenchmarkIncrementInOrDependentValues -test.trace=trace.out ./ilp/...
goos: linux
goarch: amd64
pkg: github.com/mneverov/bench/ilp
cpu: 11th Gen Intel(R) Core(TM) i7-11850H @ 2.50GHz
BenchmarkIncrementInOrDependentValues/increment_independent_values-16                  3         453596262 ns/op
BenchmarkIncrementInOrDependentValues/increment_dependent_values-16                   10         106690406 ns/op
PASS
ok      github.com/mneverov/bench/ilp   3.890s
```

Even if the assembly code is less efficient and processors cannot use ILP, incrementing dependent values now 4 times
faster 🤷.

Ok, the last attempt. Let's bench deoptimized `v1.21` version:

```sh
$ go1.21.13 test -gcflags='-N -l' -test.bench BenchmarkIncrementInOrDependentValues -test.trace=trace.out ./ilp/...
goos: linux
goarch: amd64
pkg: github.com/mneverov/bench/ilp
cpu: 11th Gen Intel(R) Core(TM) i7-11850H @ 2.50GHz
BenchmarkIncrementInOrDependentValues/increment_independent_values-16                  3         407690240 ns/op
BenchmarkIncrementInOrDependentValues/increment_dependent_values-16                    3         393600451 ns/op
PASS
ok      github.com/mneverov/bench/ilp   4.887s
```

Now incrementing dependent values is still faster and also as fast as optimized go `v1.23`.