package gdebug

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var count = 0

func DumpJSON(itf interface{}) {
	b, err := json.Marshal(itf)
	if err != nil {
		fmt.Println("Could not dumpjson:", err)
	}
	fmt.Println(string(b))
}

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

// StackIgnoreFirst retuns the stack with the first n stack items ignored
// (starting at called position)
func StackIgnoreFirst(n int) string {
	stack := strings.Split(string(debug.Stack()), "\n")
	// first line is goroutine, we always keep dat
	// then lines come in two: top line is function, bottom is file:line info

	// since we start from the second line (first is goroutine info) and since
	// we start n from the caller, we skip 2 additional (debug func + here).
	// We multiply by 2 to get rid of the pairs described above.
	lines := append(stack[0:1], stack[(n+3)*2:]...)
	return strings.Join(lines, "\n")
}

// SmartStack returns the stack starting with exactly the line that panick'ed
// (goroutine is mentioned still as first line)
func SmartStack() string {
	gopath := os.Getenv("GOPATH")
	stack := strings.Split(string(debug.Stack()), "\n")

	// we look for panic, we start at that
	count := 0
	for _, s := range stack {
		if strings.Contains(s, "runtime/panic.go:") {
			break
		}
		count++
	}

	// first line is goroutine nb, we keep dat
	lines := append(stack[0:1], stack[count+1:]...)

	var total string
	for i := range lines {
		if i == 0 {
			total += lines[i]
			continue // goroutine line
		}

		if i%2 == 1 {
			total += "\n  "
		} else {
			total += "\n"
		}
		total += strings.Replace(lines[i], gopath, "$GOPATH", 1)
	}
	return total
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
