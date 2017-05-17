package gdebug

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var count = 0

// Debugf prints the filename:line_nb of the Debugf call
func Debugf(f string, vals ...interface{}) {
	stack := strings.Split(string(debug.Stack()), "\n")
	line := stack[6]
	line = line[strings.LastIndex(line, "/")+1:]
	fmt.Printf(fmt.Sprintf("%s: ", line)+f, vals...)
}

// Dump prints a spew.Dump with filename:line_nb before
func Dump(a ...interface{}) {
	Debugf(" ")
	spew.Dump(a...)
}

// Errorf creates the error, but prepends the filename:line_nb of the creation
func Errorf(f string, vals ...interface{}) error {
	stack := strings.Split(string(debug.Stack()), "\n")
	line := stack[6]
	line = line[strings.LastIndex(line, "/")+1:]
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
