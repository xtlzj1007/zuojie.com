package main

import (
	"fmt"
	"strings"
	"bytes"
	"strconv"
)

func main() {
	var str, str1, str2 string
	fmt.Println("请输入两个整数中间使用+号，例如123+456")
	fmt.Printf(">>>") // 获取字符串
	fmt.Scanf("%s", &str)
	s := strings.Split(str, "+") // 使用+分隔数值
	str1 = s[0]
	str2 = s[1]
	sum := Add(str1, str2)
	fmt.Println(sum)
}

//计算大数之间的和
func Add(a string, b string) string {
	//位数较少的，高位补0
	if len(a) < len(b) {
		a = strings.Repeat("0", len(b)-len(a)) + a
	} else if len(a) > len(b) {
		b = strings.Repeat("0", len(a)-len(b)) + b
	}
	//fmt.Println(a,b)
	strLen := len(a)
	//保存结果的数组，从低位到高位进行保存
	nums := make([]uint8, strLen)
	//从低位到高位开始遍历a，b
	addOne := false
	for i := 0; i < strLen; i++ {
		num_a := a[strLen-i-1] - '0'
		num_b := b[strLen-i-1] - '0'
		//fmt.Println(num_a, "  ", num_b)
		sum := num_a + num_b //每一位进行加法运算
		if addOne {
			sum++
		}
		if sum > 9 {
			//进位
			sum -= 10
			addOne = true
		} else {
			addOne = false
		}
		nums[i] = sum
	}
	result := convertToString(nums)
	//进位，最前面补1
	if addOne {
		result = "1" + result
	}
	return result
}

// 将切片中的数据转换成一个字符串
func convertToString(nums []uint8) string {
	var b bytes.Buffer
	for i := len(nums) - 1; i >= 0; i-- {
		b.WriteString(strconv.Itoa(int(nums[i])))
	}
	return b.String()
}
