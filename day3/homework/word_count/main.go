package main

import (
	"fmt"
	"bufio"
	"os"
)

//获取一行可以包含空格的字符串
func Scanf(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*a = string(data)
}

//统计输入数据的英文字母，空格，数字和其他字符的个数
func main() {
	var str string
	//str = "abcd 123 ,,;;?"
	space_count := 0
	chars_count := 0
	number_count := 0
	rest_count := 0
	fmt.Println("请输入一串数据：")
	Scanf(&str)
	fmt.Println("元字符串为：", str)
	for _, v := range str {
		//fmt.Println(v)
		if uint8(v) == 32 {
			space_count ++
		} else if (65 <= uint8(v) && uint8(v) <= 90) || (97 <= uint8(v) && uint8(v) <= 122) {
			chars_count ++
		} else if 48 <= uint8(v) && uint8(v) <= 57 {
			number_count ++
		} else {
			rest_count ++
		}
	}
	fmt.Printf("空格个数:%d\n", space_count)
	fmt.Printf("数字个数:%d\n", number_count)
	fmt.Printf("英文字母个数:%d\n", chars_count)
	fmt.Printf("其他字符个数:%d\n", rest_count)
}
