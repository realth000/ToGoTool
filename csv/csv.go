package csv

func IsEqual(a [][]string, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for index, la := range a {
		if len(la) != len(b[index]) {
			return false
		}
		for i, s := range la {
			if s != b[index][i] {
				return false
			}
		}
	}
	return true
}
