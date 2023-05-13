package main

import (
	"fmt"
)

// %v - INTERPOLATE THE DEFAULT REPRESENTATION

// %s - INTERPOLATE A STRING

// %d - INTERPOLATE AN INTEGER IN DECIMAL FORM

// %f - INTERPOLATE A DECIMAL; %.2F - ROUNDED DECIMAL

func main () {

	// fmt.Printf("I'm %v years old\n", 43)

	// fmt.Printf("I'm %v years old\n", "way too many")

	// fmt.Printf("I'm %s years old\n", "way too many")

	// fmt.Printf("I'm %d years old\n", 43)

	// fmt.Printf("I'm %f years old\n", 43.523)
	// fmt.Printf("I'm %.2f years old\n", 43.523)

	// Task 1
	const name = "Saul Goldman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)

	fmt.Println(msg)

	

}