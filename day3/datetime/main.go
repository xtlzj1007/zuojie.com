package main

import (
	"time"
	"fmt"
)

func PringTime() {
	now := time.Now()
	fmt.Printf("type of now is %T\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	fmt.Printf("%04d-%02d-%02d %02d:%02d:%02d\n", year, month, day, now.Hour(), now.Minute(), now.Second())
	fmt.Println(now.Unix())
	str := now.Format("2006/01/02 15:04:05")
	fmt.Println(str)
}

func testStrReverseUtf8() {
	var str = "abcdef世界您好"
	//str1 := str[6:]
	var b = []rune(str)
	for i := 0; i < 1000000; i++ {
		for i := 0; i < len(b)/2; i++ {
			b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
		}
		//fmt.Printf("str1:%s\n", string(b))
		//fmt.Println(len(b))
		//time.Sleep(1)
	}
}
func printRunTime() {
	start := time.Now().UnixNano()
	testStrReverseUtf8()
	end := time.Now().UnixNano()
	fmt.Printf("%v", (end-start)/1000)
}

func main() {
	//PringTime()
	printRunTime()
}
