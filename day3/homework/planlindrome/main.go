package main

import (
	"fmt"
)

//判断输入的是否是回文

func main() {
	var str string
	fmt.Println("请输入一串字符串：")
	fmt.Scanf("%s", &str)
	str1 := []rune(str)
	for i := 0; i < len(str1)/2; i++ {
		str1[i], str1[len(str1)-i-1] = str1[len(str1)-i-1], str1[i]
	}
	str2 := string(str1)
	if str == str2 {
		fmt.Printf("%s是回文！", str)
	} else {
		fmt.Printf("%s不是回文！", str)
	}
}
