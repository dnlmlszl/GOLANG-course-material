package main

import (
	"fmt"
)

func getLength (email string) int {
	
	return len(email)

}

func main () {

	// Task 1
	// messageLen := 10
	// maxMessageLen := 20

	// fmt.Println("Tryingto send a message of length:", messageLen)

	// if messageLen <= maxMessageLen {
	// 	fmt.Println("Message sent")
	// } else if messageLen == maxMessageLen {
	// 	fmt.Println("Mukulu buki")
	// } else {
	// 	fmt.Println("Message not sent")
	// }

	// Task 2
	email := "blablabla"
	
	if length := getLength(email); length < 1 {
		fmt.Println("Email is invalid")
	} else {
		fmt.Println("Email is valid", length)
	}
}