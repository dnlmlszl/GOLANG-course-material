package main


import (
	"fmt"
	// "errors"
)
// Declare a func
func sub(x int, y int) int {
	return x - y
}
// Also convenient way to declare
func concat(s1, s2 string) string {
	return s1 + s2
}
// Named return
func getCoords() (z, k int) {
	return z, k
}

func main () {
	// Named return
	test(17)
	test(22)
	test(20)
	
	// Ignoring return values
	firstName, _ := getNames()
	fmt.Println("Welcome to textio", firstName)

	// Increment
	x := 5
	increment(&x)
	fmt.Println(x)

	sendsSofar := 430
	const sendsToAdd = 25
	
	sendsSofar = incrementSends(sendsSofar, sendsToAdd)
	fmt.Println("You've sent", sendsSofar, "messages")

	// Basic funcs
	fmt.Println(sub(7, 8))
	fmt.Println(concat("Lane", " happy birthday"))
	fmt.Println(concat("Go", " is fantastic"))
}

func incrementSends(sendsSofar, sendsToAdd int) int {
	sendsSofar = sendsSofar + sendsToAdd

	return sendsSofar
}

func increment(x *int) {
	(*x)++
}

func getNames() (string, string) {
	return "John", "Doe"
}

func yearsUntilEvents(age int) (yearsUntilAdult int, yearsUntilDrinking int, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age

	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}

	yearsUntilDrinking = 21 - age

	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}

	yearsUntilCarRental = 25 - age

	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}

func test(age int) {
	fmt.Println("Age:", age, "years")
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
	fmt.Println("You are an adult in", yearsUntilAdult, "years")
	fmt.Println("You can drink in", yearsUntilDrinking, "years")
	fmt.Println("You can rent a car in", yearsUntilCarRental, "years")
}

// Early returns
// func divide(dividend, divisor int) (int, error) {
// 	if divisor == 0 {
// 		return 0, errors.New("Can't divide by zero")
// 	}
// 	return dividend / divisor, nil
// }

// func getInsuranceAmount(status insuranceStatus) int {
// 	amount := 0
// 	if !status.hasInsurance() {
// 		amount = 1
// 	} else {
// 		if status.isTotaled() {
// 			amount = 10000
// 		} else {
// 			if status.isDented() {
// 				amount = 160
// 				if status.isBigDent() {
// 					amount = 270
// 				}
// 			} else {
// 				amount = 0
// 			}
// 		}
// 	}
// 	return amount
// }

// func getInsuranceAmount(status insuranceStatus) int {
// 	amount := 0
// 	if !status.hasInsurance() {
// 		amount = 1
// 	}

// 	if status.isTotaled() {
// 		amount = 10000
// 	}
	
// 	if status.isDented() {
// 		amount = 160
// 	}

// 	if status.isBigDent() {
// 		amount = 270
// 	}
// 	return amount
// }