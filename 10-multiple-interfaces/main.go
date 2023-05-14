package main

import (
	"fmt"
)

type expense interface {
	cost() float64
}

type printer interface {
	print() 
}

type email struct {
	isSubscribed bool
	body string
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	}
	return 0.01 * float64(len(e.body))
}

func (e email) print() {
	fmt.Println(e.body)
}

func print(p printer) {
	p.print()
}

func test(e expense, p printer) {
	fmt.Printf("Printing with cost $%.2f ...\n", e.cost())
	p.print()
}


func main () {
	e := email {
		isSubscribed: true,
		body: "Hello Kótyonfitty",
	}
	test(e, e)
	e = email {
		isSubscribed: false,
		body: "I want my money back, Mr. Kótyonfitty",
	}
	test(e, e)
	e = email {
		isSubscribed: true,
		body: "I want more content for the money",
	}
	test(e, e)
	e = email {
		isSubscribed: false,
		body: "I want more content for the money",
	}
	test(e, e)
}

// Example naming
// type Copier interface {
// 	Copy(string, string) int
// }

// type CopierII interface	{
// 	CopyII(sourceFile string, destinationFile string) (bytesCopied int)
// }