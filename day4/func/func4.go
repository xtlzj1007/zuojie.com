package main

import (
	"fmt"
)

func add(a, b int) (c int) {
	c = a + b
	return
}

func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = sum / 2
	return
}

func Add(arg ... int) int {
	var sum int
	for i := 0; i < len(arg); i++ {
		sum = sum + arg[i]
	}
	return sum
}

func main() {
	sum := add(100, 200)
	fmt.Println(sum)
	sum, avg := calc(100, 200)
	fmt.Println(sum, avg)
	fmt.Println(Add(12, 34, 5, 6))
}



