package main

import (
	"fmt"
)

func main() {
	Pulsn(5)
}

// 循环添加
func Pulsn(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("%d+%d=%d\n", i, n-i, n)
	}
}
