package main

import (
	"fmt"
	"strconv"
)

func main() {
	// for i := 0; i < 10; i++ {

	// fmt.Println(" "+strconv.Itoa(i)+" Loops")
	// }

	numbers := []int{1, 2, 3, 4}
	for _, value := range numbers {
		fmt.Println("Number " + strconv.Itoa(value))
	}

	for i := 1; i <= 5; i++ {
		fmt.Println("Number " + strconv.Itoa(i))
	}
	fmt.Println(add(13, 14))
}

func add(x int, y int) int {

	return x + y
}
