package gdebug

import "fmt"

var count = 0

// CountPrintf prints the debug message along with a prefix number to identify it
func CountPrintf(f string, vals ...interface{}) {
	count++
	fmt.Println("%d - "+f, count, vals)
}

// ResetCount resets the count used in CountPrinf
func ResetCount() {
	count = 0
}
