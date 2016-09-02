package gstring

// SetIfExists does dst = src (only if src is not empty)
func SetIfExists(dst *string, src string) {
	if src != "" {
		*dst = src
	}
}

// MatchingProblem returns a formatted string which explains
// that the two elements s1 and s2 of type obj don't match
func MatchingProblem(obj, s1, s2 string) string {
	return obj + " " + s1 + " and " + s2 + " don't match"
}
