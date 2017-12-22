package fibo

import "fmt"

func PrintFibo(n int) {
	var a int
	var b int
	a, b = 0, 1
	for i := 1; i < n; i++ {
		fmt.Printf("%d\t", a)
		a, b = b, a+b
	}
}
