package main

import (
	"fmt"
	"time"
)

func fibo(pipe chan uint64, n int) {
	var a uint64
	var b uint64
	a, b = 0, 1
	for i := 0; i < n; i++ {
		//fmt.Println(a)
		pipe <- a
		a, b = b, a+b
	}
}

func main() {
	n := 100
	pipe := make(chan uint64, 10)
	go fibo(pipe, n)
	time.Sleep(2 * time.Second)
	for i := 0; i < n; i++ {
		fib := <-pipe
		fmt.Println(fib)
	}
}
