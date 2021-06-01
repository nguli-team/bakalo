package util

// FindString returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func FindString(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

// ContainsString tells whether a contains x.
func ContainsString(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
