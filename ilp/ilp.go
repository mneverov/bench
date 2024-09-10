package ilp

func IncrementDependentValues() (int, int) {
	numIter := 256 * 1024 * 1024
	a := [2]int{}

	for i := 0; i < numIter; i++ {
		a[0]++
		a[0]++
	}

	return a[0], a[1]
}

func IncrementIndependentValues() (int, int) {
	numIter := 256 * 1024 * 1024
	a := [2]int{}

	for i := 0; i < numIter; i++ {
		a[0]++
		a[1]++
	}
	return a[0], a[1]
}
