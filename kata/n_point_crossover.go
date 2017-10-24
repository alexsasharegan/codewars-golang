package kata

// Crossover kata.
// ns : slice of indices
// xs, ys : chromosomes as slices of ints
func Crossover(ns, xs, ys []int) ([]int, []int) {
	visited := make(map[int]bool)
	xs1 := make([]int, len(xs))
	ys1 := make([]int, len(ys))
	copy(xs1, xs)
	copy(ys1, ys)

	// we don't need the index of the index, just its value
	for _, i := range ns {
		if !visited[i] {
			visited[i] = true
			xs1 = append(xs1[:i], ys[i:]...)
			ys1 = append(ys1[:i], xs[i:]...)
			copy(xs, xs1)
			copy(ys, ys1)
		}
	}

	return xs1, ys1
}
