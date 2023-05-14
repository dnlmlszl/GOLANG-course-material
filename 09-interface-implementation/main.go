package main

import (
	"fmt"
)

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name string
	hourlyPay int
	hoursPerYear int
}

type fullTime struct {
	name string
	salary int 
}

func (c contractor) getName() string {
	return c.name 
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}


func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func test(e employee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("====================================")
}


func main () {
	test(fullTime {
		name: "Jonathan T. Dobson",
		salary: 68000,
	})
	test(contractor{
		name: "Edyta C. Kowalsky",
		hourlyPay: 20,
		hoursPerYear: 2650,
	})
	test(fullTime {
		name: "Kalman T. Szörnyeteg",
		salary: 72000,
	})
	test(contractor{
		name: "Erwin V. Rommel",
		hourlyPay: 24,
		hoursPerYear: 3200,
	})
}