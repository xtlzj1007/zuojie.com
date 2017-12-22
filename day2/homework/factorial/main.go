package main

import "fmt"

// 阶乘程序
func main() {
	fmt.Println(PlusFactorial(10000))
}

// 计算阶乘的函数
func Factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}

// 计算阶乘之间的加法
func PlusFactorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return Factorial(n) + PlusFactorial(n-1)
	}
}
