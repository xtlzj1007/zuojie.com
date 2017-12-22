package main

import ("fmt"
	   "zuojie.com/day1/homework"
)


func main(){
	pipe := make(chan string,1)
	pipe <- "Hello World"   // 数据量超过管道maxsize的时候，死锁检测
	var a string
	a = <- pipe // (=<-) get 从管道里取出来
	fmt.Printf("%s\n",a)
	homework.PrintData()
}

