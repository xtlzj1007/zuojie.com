package main

import "fmt"

func testDefer() {
	var a int = 0
	defer fmt.Println(a)
	a ++
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}



func main() {
	//testDefer()
	f()
}

