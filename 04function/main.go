package main

import "fmt"

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

func main() {
	a, b := 5, 4
	result1, result2 := calc2(a, b)

	fmt.Println(result1, result2)
}
