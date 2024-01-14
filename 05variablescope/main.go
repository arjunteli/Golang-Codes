package main

import (
	"fmt"
	"time"
)

var a = 10

func demo() {
	a = 100
	fmt.Println("The global", a)
	//return
}

func main() {
	fmt.Println("Variable Scope")
	demo()
	fmt.Println(a)
	fmt.Println("You entered", a, "which is a global", a)
	if a > 90 {
		fmt.Println(a, "is greater than 90")
	} else if a > 90 {
		fmt.Println(a, "is less than 90")
	} else if a < 89 {

	} else {

	}

	switch a {
	case 100:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Default")
	}
	fmt.Println(time.Now().Weekday())
}
