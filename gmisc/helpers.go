package gmisc

import "fmt"

// PanicIf panics if the interface passed is not nil
func PanicIf(i interface{}) {
	if i != nil {
		panic(i)
	}
}

// Panicf panics with the format string and its parameters
func Panicf(format string, a ...interface{}) {
	panic(fmt.Sprintf(format, a))
}
