package gdebug

import "fmt"

var count = 0

// CountPrintf prints the debug message along with a prefix number to identify it
func CountPrintf(f string, vals ...interface{}) {
	count++
	fmt.Printf(fmt.Sprintf("%d - ", count)+f, vals...)
}

// CountPrintln prints the debug message along with a prefix number to identify it
func CountPrintln(f ...interface{}) {
	count++
	fmt.Printf("%d - ", count)
	fmt.Println(f...)
}

// ResetCount resets the count used in CountPrinf
func ResetCount() {
	count = 0
}
