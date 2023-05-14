package main

import (
	"fmt"
)

func printNumericValues(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func main () {
	printNumericValues(1)
	printNumericValues("1")
	printNumericValues(struct{}{})
}

// CLEAN INTERFACE
// type File interface {
// 	io.Closer
// 	io.Reader
// 	io.Seeker
// 	Readdir(count int) ([]os.FileInfo, error)
// 	Stat() (os.FileInfo, error)
// }