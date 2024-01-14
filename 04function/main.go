package main

import (
	"fmt"
	"strconv"
)

//const loginToken = "Login Value"

func calc1(x, y int) (int, int) {
	var out1 = x + y
	var out2 = x - y
	return out1, out2
}

func calc2(x, y int) (out1, out2 int) {
	out1 = x + y
	out2 = x - y
	return
}
func printTable(num, i int) {
	if i == 11 {
		return
	}
	fmt.Println(strconv.Itoa(num) + " x " + strconv.Itoa(i) + " = " + strconv.Itoa(num*i))
	printTable(num, i+1)
}

func main() {
	// a, b := 5, 4
	// result1, result2 := calc2(a, b)

	// fmt.Println(result1, result2)

	num := 2
	printTable(num, 1)
}
