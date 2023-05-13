package main

import (
	"fmt"
)

func sub(x int, y int) int {
	return x - y
}

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func main () {
	fmt.Println(sub(7, 8))
	fmt.Println(concat("Lane", " happy birthday"))
	fmt.Println(concat("Go", " is fantastic"))
}