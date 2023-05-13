package main

import (
	"fmt"
)

func main () {
	// Task 1
	// fmt.Println("hello world")

	// Task 2
	// messagesFromDoris := [] string {
	// 	"You doing anything later?",
	// 	"Did you get my last message?",
	// 	"Don't leave me hanging...",
	// 	"Please respond I'm lonely!",
	// }
	// numMessages := float64(len(messagesFromDoris))
	// costPerMessage := 0.02

	// totalCost := numMessages * costPerMessage

	// fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)

	// Task 3
	var username string = "wagslane"
	var password string = "12345678912345"

	fmt.Println("Authorization: Basic", username + ": " + password)
}