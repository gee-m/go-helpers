package gstrings

// Find returns the index of the needle in the haystack (has to match exactly)
// or returns -1 if the needles wasn't found
func Find(needle string, haystack []string) int {
	for i := range haystack {
		if needle == haystack[i] {
			return i
		}
	}
	return -1
}

// FindCust returns the idnex of the needle in the haystack (has to match according
// to the custCompare function) or returns -1 if the needles wasn't found
func FindCust(needle string, haystack []string, custCompare func(string, string) bool) int {
	for i := range haystack {
		if custCompare(needle, haystack[i]) {
			return i
		}
	}
	return -1
}
