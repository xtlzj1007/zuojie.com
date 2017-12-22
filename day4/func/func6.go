package main

import "fmt"

func calc(n int) int {
	if n == 1 {
		return 1
	}
	return calc(n-1) * n
}

func fibo(n int) int {
	if n <= 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

func main() {
	//n := calc(5)
	//fmt.Println(n)
	for i := 0; i < 5; i++ {
		n := fibo(i)
		fmt.Println(n)
	}

}
