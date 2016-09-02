package gmisc

func PanicIf(i interface{}) {
	if i != nil {
		panic(i)
	}
}
