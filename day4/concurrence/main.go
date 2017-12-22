package main

import (
	"fmt"
)

//func Add(x, y int) {
//	z := x + y
//	fmt.Println(z)
//}

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func main() {
	chs := make([] chan int, 10)
	for i := 0; i < 10; i++ {
		//go Add(i, i)
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	//time.Sleep(time.Second * 1)
	for _, ch := range (chs) {
		<-ch
	}
}
