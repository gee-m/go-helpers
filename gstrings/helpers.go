package gstrings

import (
	"fmt"
	"strings"
	"unicode"
)

// SetIfExists does dst = src (only if src is not empty)
func SetIfExists(dst *string, src string) {
	if src != "" {
		*dst = src
	}
}

// MatchingProblem returns a formatted string which explains
// that the two elements s1 and s2 of type obj don't match
func MatchingProblem(obj, i1, i2 interface{}) string {
	return fmt.Sprintf("%s %s and %s do not match", obj, i1, i2)
}

// ToCamel formats string s into camel case from snake case
// with specified rune
// e.g. hi_im_camel -> hiImCamel
func ToCamel(s string, sep rune, capitalize bool) string {
	s = strings.ToLower(s)

	cpy := s
	found := 0
	for i, v := range cpy {
		first := i == 0
		last := i == len(cpy)-1
		if first && capitalize {
			s = strings.ToUpper(s[:1]) + s[1:]
		}
		if first && v == sep {
			continue
		} else if !first && cpy[i-1] == byte(sep) && v == sep {
			continue
		} else if !last && v == sep && cpy[i+1] == byte(sep) {
			continue
		}

		if v == sep && !last {
			s = s[:i-found] + strings.ToUpper(cpy[i+1:i+2]) + cpy[i+2:]
			found++
		}
	}
	return s
}

// ToSnake formats string s into snake_case from CamelCase
func ToSnake(s string) string {
	cpy := s
	found := 0
	lastTwoCapitalized := false

	for i, v := range cpy {
		first := i == 0

		if first { // skip if first letter
			continue
		}

		lastCapitalized := unicode.IsUpper(rune(cpy[i-1]))
		currCapitalized := unicode.IsUpper(v)

		if lastCapitalized && currCapitalized {
			lastTwoCapitalized = true
			continue
		}
		if (!lastCapitalized && currCapitalized) || (lastTwoCapitalized && !currCapitalized) {
			// We want to separate with underscore
			s = s[:i+found] + "_" + s[i+found:]
			found++ // shift the index of the string s
		}
		lastTwoCapitalized = false
	}
	return strings.ToLower(s)
}

// StringBetween returns a string enclosed by start and end (excluded both)
func StringBetween(s, start, end string) (found bool, str string) {
	i1 := strings.Index(s, start)
	if i1 == -1 {
		return false, ""
	}
	s2 := s

	if start != "" {
		s2 = s[i1+len(start):]
		if end == "" {
			return true, s2
		}
	}

	i2 := strings.Index(s2, end)
	if i2 == -1 {
		return false, ""
	}
	return true, s2[:i2]
}
