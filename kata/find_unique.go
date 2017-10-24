package kata

// FindUniq kata.
func FindUniq(arr []float32) float32 {
	// First item is unique
	if arr[0] != arr[1] && arr[1] == arr[2] {
		return arr[0]
	}
	// Index of subslice == i-1 of arr,
	// effectively compares element w/ predecessor.
	for i, v := range arr[1:] {
		if v != arr[i] {
			return v
		}
	}
	return 0
}
