package main

import (
	"fmt"
)

type messageToSend struct {
	phoneNumber int
	message string
}

type newMessageToSend struct {
	msg string
	sender user
	recipient user
}

type user struct {
	name string
	number int
}

type car struct {
	Company string
	Modell string
	Height int
	Width int
	Wheel struct {
		Radius int
		Material string
	}
}

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("===============================")
}

func canSendMessage(mToSend newMessageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}
	return true
}

func test2(mToSend newMessageToSend) bool {
	if mToSend.sender.name == "" || mToSend.recipient.name == "" || mToSend.sender.number == 0 || mToSend.recipient.number == 0  {
		fmt.Printf("Cannot send message: %v (%v) to %v (%v)\n", mToSend.sender.name, mToSend.sender.number, mToSend.recipient.name, mToSend.recipient.number)
		return canSendMessage(mToSend)
	}
	fmt.Printf("Sending message: '%s' from: %v (%v) to %v (%v)\n", mToSend.msg, mToSend.sender.name, mToSend.sender.number, mToSend.recipient.name, mToSend.recipient.number)
	fmt.Println("===============================")
	return canSendMessage(mToSend)
	
}

func main () {
	myCar := struct {
		Make string
		Model string
	} {
		Make: "Tesla",
		Model: "Model 3",
	}

	fmt.Printf("%v %v\n", myCar.Model, myCar.Make)
	fmt.Println("===============================")

	newCar := car{
		Company: "BMW",
		Modell:  "Porördög 3",
		Height:  150,
		Width:   200,
		Wheel: struct {
			Radius   int
			Material string
		}{
			Radius:   15,
			Material: "Alloy",
		},
	}

	fmt.Printf("%v %v\n", newCar.Modell, newCar.Wheel.Material)
	fmt.Println("===============================")

	test(messageToSend {
		phoneNumber: 148255510981,
		message: "Thanks for hit me up",
	})
	test(messageToSend {
		phoneNumber: 548255510981,
		message: "Thanks for cannot hit me up",
	})
	test(messageToSend {
		phoneNumber: 648255510981,
		message: "Thanks for do not hit me up",
	})
	test2(newMessageToSend {
		msg: "You have an appointment",
		sender: user {
			name: "John Remember",
			number: 12345678,
		},
		recipient: user {
			name: "Clara Halafax",
			number: 12345678,
		},
	})
	test2(newMessageToSend {
		msg: "You have an appointment",
		sender: user {
			name: "",
			number: 12345678,
		},
		recipient: user {
			name: "Clara Halafax",
			number: 0,
		},
	})
	test2(newMessageToSend {
		msg: "You have an appointment",
		sender: user {
			name: "John Remember",
			number: 12345678,
		},
		recipient: user {
			name: "",
			number: 0,
		},
	})
	
}