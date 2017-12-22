package main

import "fmt"
import (
	//_ "zuojie.com/day2/init_func"
	"zuojie.com/day2/swap"
)

//const (
//	a = 1 << iota
//	b
//	c
//)

func init() {
	fmt.Println("init 执行了！")
}

func main() {
	a := 100
	b := 200
	a, b = swap.Swap(a, b)
	fmt.Printf("a=%d,b=%d", a, b)
	//fmt.Println("main 执行了！")
	//fmt.Println(c)

}
