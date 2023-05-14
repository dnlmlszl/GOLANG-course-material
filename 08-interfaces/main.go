package main

import (
	"fmt"
	"math"
	"time"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2 * r.width + 2 * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * c.radius * math.Pi
}

///////////////////////////////////////////////////////////////////////
func sendMessage (msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime time.Time
	recipientName string
}

type sendingReport struct {
	reportName string
	numberOfSends int
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birtday on %s\n", bm.recipientName, bm.birthdayTime)
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf("Your %s is ready. You've sent %v messages\n", sr.reportName, sr.numberOfSends)
}


func test(m message) {
	sendMessage(m)
	fmt.Println("====================================")
}

func main () {
	r := rect {
		width: 5.0,
		height: 10.0,
	}

	c := circle {
		radius: 10.0,
	}

	fmt.Println(r.area(), r.perimeter(), c.area(), c.perimeter())

	//Testcases for message interface
	test(sendingReport{
		reportName: "Analysis for animal habits",
		numberOfSends: 15,
	})
	test(birthdayMessage{
		recipientName: "Daniel.M.Laszlo",
		birthdayTime: time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
	test(sendingReport{
		reportName: "Gaming for animal habits",
		numberOfSends: 11,
	})
	test(birthdayMessage{
		recipientName: "Sarah.M.Connor",
		birthdayTime: time.Date(1999, 07, 21, 0, 0, 0, 0, time.UTC),
	})

}