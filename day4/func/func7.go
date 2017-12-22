package main

import "fmt"

//闭包函数，返回值是一个函数
func Adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func main() {
	f := Adder()
	fmt.Println(f(1))
	fmt.Println(f(20))
	fmt.Println(f(300))
}


