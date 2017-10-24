package kata

// NbMonths kata.
func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	const dr = 0.5
	var s float64
	var m int
	old := float64(startPriceOld)
	new := float64(startPriceNew)

	for s+old < new {
		m++
		if m%2 == 0 {
			percentLossByMonth += dr
		}
		s += float64(savingperMonth)
		old -= old * percentLossByMonth / 100
		new -= new * percentLossByMonth / 100
	}

	return [2]int{m, roundToInt(s + old - new)}
}

func roundToInt(a float64) int {
	if a < 0 {
		return int(a - 0.5)
	}
	return int(a + 0.5)
}
