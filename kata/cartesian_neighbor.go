package kata

// CartesianNeighbor kata.
func CartesianNeighbor(x, y int) [][]int {
	n := make([][]int, 0, 8)

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			n = append(n, []int{i, j})
		}
	}

	return n
}
