package main

import "fmt"

/*
九九乘法表
 */
func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dX%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
