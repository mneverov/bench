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

func Increment4DependentValues() (int, int, int, int) {
	numIter := 256 * 1024 * 1024
	a := [4]int{}

	for i := 0; i < numIter; i++ {
		a[0]++
		a[0]++
		a[0]++
		a[0]++
	}

	return a[0], a[1], a[2], a[3]
}

func Increment4IndependentValues() (int, int, int, int) {
	numIter := 256 * 1024 * 1024
	a := [4]int{}

	for i := 0; i < numIter; i++ {
		a[0]++
		a[1]++
		a[2]++
		a[3]++
	}
	return a[0], a[1], a[2], a[3]
}

func Increment8DependentValues() (int, int, int, int, int, int, int, int) {
	numIter := 256 * 1024 * 1024
	a := [8]int{}

	for i := 0; i < numIter; i++ {
		a[0]++
		a[0]++
		a[0]++
		a[0]++
		a[0]++
		a[0]++
		a[0]++
		a[0]++
	}

	return a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]
}

func Increment8IndependentValues() (int, int, int, int, int, int, int, int) {
	numIter := 256 * 1024 * 1024
	a := [8]int{}

	for i := 0; i < numIter; i++ {
		a[0]++
		a[1]++
		a[2]++
		a[3]++
		a[4]++
		a[5]++
		a[6]++
		a[7]++
	}
	return a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]
}
