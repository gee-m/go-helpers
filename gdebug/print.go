package gdebug

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var count = 0

// LineInfo returns "{nameoffile}.go:{linenumber}" of the caller
func LineInfo() string {
	return lineInfo(8)
}

// lineInfo returns "{nameoffile}.go:{linenumber}" of the caller, offsetted by n
func lineInfo(n int) string {
	stack := strings.Split(string(debug.Stack()), "\n")
	line := stack[n]
	line = line[strings.LastIndex(line, "/")+1:]
	line = strings.Split(line, " ")[0]
	return line
}

// Debugf prints the filename:line_nb of the Debugf call
func Debugf(f string, vals ...interface{}) {
	line := lineInfo(8)
	fmt.Printf(fmt.Sprintf("%s: ", line)+f, vals...)
}

// Dump prints a spew.Dump with filename:line_nb before
func Dump(a ...interface{}) {
	line := lineInfo(8)
	fmt.Printf("%s: ", line)
	spew.Dump(a...)
}

// Errorf creates the error, but prepends the filename:line_nb of the creation
func Errorf(f string, vals ...interface{}) error {
	line := lineInfo(8)
	return fmt.Errorf(fmt.Sprintf("%s: ", line)+f, vals...)
}

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
