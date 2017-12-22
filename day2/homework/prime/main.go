package main

import (
	"fmt"
)

//计算素数
func PrintPrime(n int, m int) {
	count := 0
	for i := n; i < m; i++ {
		for j := 2; j <= i; j++ {
			if i%j == 0 {
				if i == j {
					fmt.Printf("%d\t", i)
					count ++
				} else {
					break
				}
			}
		}
	}
	fmt.Printf("\n%d~%d之间素数的个数为：%d", n, m, count)
}

func main() {
	PrintPrime(1, 200000)
}
