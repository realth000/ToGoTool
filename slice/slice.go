package slice

// CleanDuplicate removes duplicate items in slice.
func CleanDuplicate[T comparable](slice []T) []T {
	tmp := make(map[T]bool)
	var ret []T
	for _, v := range slice {
		if _, ok := tmp[v]; !ok {
			tmp[v] = true
			ret = append(ret, v)
		}
	}
	return ret
}
