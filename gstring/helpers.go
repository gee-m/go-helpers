package gstring

import "fmt"

// SetIfExists does dst = src (only if src is not empty)
func SetIfExists(dst *string, src string) {
	if src != "" {
		*dst = src
	}
}

// MatchingProblem returns a formatted string which explains
// that the two elements s1 and s2 of type obj don't match
func MatchingProblem(obj, i1, i2 interface{}) string {
	return fmt.Sprintf("%s %s an %s do not match", obj, i1, i2)
}
