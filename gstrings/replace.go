package gstrings

// ReplaceFunc is the type of a function which takes a string and replaces
// what it wants, then returns it
type ReplaceFunc func(string) string

// ReplaceFuncBetween is similar to ReplaceAnyBetween.
// It looks through the string, and replaces the string between start and end
// with the ReplaceFunc.
func ReplaceFuncBetween(s, start, end string, rpl ReplaceFunc) string {
	var newS string
	startLen := len(start)
	endLen := len(end)

	var (
		in      int
		inStart int
	)

	for i := 0; i < len(s); i++ {
		if (in == 0 && i+startLen >= len(s)) || (in != 0 && i+endLen > len(s)) {
			newS += s[i:]
			break
		}
		if s[i:i+startLen] == start {
			if in == 0 {
				inStart = i + startLen
				newS += start
			}
			i += startLen - 1
			in++
			continue
		} else if in != 0 && s[i:i+endLen] == end {
			in--
			if in == 0 {
				newS += rpl(s[inStart:i]) + end
				inStart = 0
			}
			i += endLen - 1
			continue
		}
		if in == 0 {
			newS += string(s[i])
		}
	}

	return newS
}

// ReplaceAnyBetween replaces in s only 'old' between start and end for 'new'
// e.g. "asdf{{hehe}}ff" -(start:}},end:}},old:ha,new:ho)-> "asdf{{hoho}}ff"
func ReplaceAnyBetween(s, start, end, old, new string) string {
	var newS string
	startLen := len(start)
	endLen := len(end)
	oldLen := len(old)

	var in int

	for i := 0; i < len(s); i++ {
		if (in == 0 && i+startLen >= len(s)) || (in != 0 && i+endLen > len(s)) {
			newS += s[i:]
			break
		}
		if s[i:i+startLen] == start {
			newS += s[i : i+startLen]
			i += startLen - 1
			in++
			continue
		} else if in != 0 && s[i:i+endLen] == end {
			newS += s[i : i+endLen]
			i += endLen - 1
			in--
			continue
		}

		if in != 0 && s[i:i+oldLen] == old {
			newS += new
			i += oldLen - 1
		} else {
			newS += string(s[i])
		}
	}

	return newS
}
