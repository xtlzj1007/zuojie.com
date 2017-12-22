package main

import "fmt"

func PrintData(){
	num       := 1024
	str       := "Hello World!"
	float_num := 3.14
	bin_num   := 8
	hex_num   := 10
	fmt.Printf("整数：%d\n浮点数：%.2f\n字符串：%s\n二进制：%b\n十六进制：%x\n",num,float_num,str,bin_num,hex_num)
}

func  main() {
	PrintData()
}