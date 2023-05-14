package main

import (
	"fmt"
)

type authenticationInfo struct {
	username string
	password string
}

type rect struct {
	width int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (a authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s %s", a.username, a.password)
}

func test(authInfo authenticationInfo) {
	if authInfo.username == "" || authInfo.password == "" {
		fmt.Println("No valid information")
		return
	}
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("========================")
}

func main () {
	
	r := rect {
		width: 5,
		height: 10,
	}

	fmt.Println(r.area())

	test(authenticationInfo{
		username: "Daniel",
		password: "12345678",
	})
	test(authenticationInfo{
		username: "Kati",
		password: "12345678",
	})
	test(authenticationInfo{
		username: "Sanyi",
		password: "12345678",
	})
	test(authenticationInfo{
		username: "",
		password: "12345678",
	})
	test(authenticationInfo{
		username: "Sanyi",
		password: "",
	})
}