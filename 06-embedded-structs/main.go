package main

import (
	"fmt"
)

type car struct {
	make string
	model string
}
type truck struct {
	car 
	bedSize int
}

type sender struct {
	rateLimit int
	user
}

type user struct {
	name string
	number int
}

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rate limit:", s.rateLimit)
	fmt.Println("=================================")
}

func main () {
	lanesTruck := truck {
		bedSize: 10,
		car: car {
			make: "Toyota",
			model: "Camry",
		},
	}

	fmt.Println(lanesTruck.bedSize)
	fmt.Println(lanesTruck.make)
	fmt.Println(lanesTruck.model)

	test(sender {
		rateLimit: 2500,
		user: user {
			name: "Totya bá'",
			number: 1234587,
		},
	})
	test(sender {
		rateLimit: 2800,
		user: user {
			name: "Feri bá'",
			number: 12634587,
		},
	})
	test(sender {
		rateLimit: 2500,
		user: user {
			name: "Sally néne",
			number: 1234587,
		},
	})
}