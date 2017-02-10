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

// CamelToSnake formats string s into snake case from camel case
// e.g. hi_im_camel -> hiImCamel
func CamelToSnake(s string, capitalize bool) string {
	cpy := s
	found := 0
	for i, v := range cpy {
		first := i == 0
		last := i == len(cpy)-1
		if first && capitalize {
			s = strings.ToUpper(s[:1]) + s[1:]
		}
		if first && v == '_' {
			continue
		} else if !first && cpy[i-1] == '_' && v == '_' {
			continue
		} else if !last && v == '_' && cpy[i+1] == '_' {
			continue
		}

		if v == '_' && !last {
			s = s[:i-found] + strings.ToUpper(cpy[i+1:i+2]) + cpy[i+2:]
			found++
		}
	}
	return s
}

// SnakeToCamel formats string s into camel case from snake case
// e.g. HiImSnake -> hi_im_snake
func SnakeToCamel(s string) string {
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

// StringBetween returns a string enclosed by start and end (exluded both)
func StringBetween(s, start, end string) string {
	i1 := strings.Index(s, start)
	if i1 == -1 {
		return ""
	}
	s2 := s[i1+len(start):]
	i2 := strings.Index(s2, end)
	if i2 == -1 {
		return ""
	}
	return s2[:i2]
}
