package main

import (
	"math/rand"
	"fmt"
)

func genRandStr() {
	var a [100]string
	var b string = "1234567890abcdef"
	//var runeArr = [] rune
	for i := 0; i < len(a); i++ {
		var str string
		for j := 0; j < 32; j++ {
			// 产生随机数下标
			index := rand.Intn(len(b))
			// 格式化并返回一个字符串
			str = fmt.Sprintf("%s%c", str, b[index])
		}
		a[i] = str
	}
	for i := 0; i < 100; i++ {
		fmt.Println(a[i])
	}
}

func genSection() {
	var a [5]int
	var b = a[1:3]
	b[0] = 100
	b[1] = 200
	fmt.Printf("b:%#v\n", b)
}

func Sum(a [] int) int {
	sum := 0
	//for i := 0; i < len(a); i++ {
	//	sum += a[i]
	//}
	for _, v := range a {
		sum += v
	}
	return sum
}

func testSliceCap() {
	a := make([]int, 5, 10)
	a[4] = 100
	b := a[1:10]
	b[8] = 10
	fmt.Printf("a=%#v,cap(a)=%d,len(a)=%d\n", a, cap(a), len(a))
	fmt.Printf("b=%#v,cap(b)=%d,len(b)=%d\n", b, cap(b), len(b))

}

func testCopy() {
	s1 := [] int{1, 2, 3}
	s2 := make([]int, 10)
	copy(s2, s1)
	fmt.Println(s2)
}

func testAppend() {
	s1 := [] int{}
	//s1 = make([]int, 5)
	var b []int = []int{6, 7, 8}
	s1 = append(s1, b...) // 在之后添加{0,0,0,0,0,1,2,3}
	//s2 := s1[3:3]
	fmt.Printf("a=%#v", s1)
}

func testStrReverseUtf8() {
	var str = "abcdef世界您好"
	//str1 := str[6:]
	var b = []rune(str)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	fmt.Printf("str1:%s\n", string(b))
	fmt.Println(len(b))
}



func main() {
	//genRandStr()
	//genSection()
	//a := [200] int{}
	//for i := 0; i < 200; i++ {
	//	a[i] = i
	//}
	//sum := Sum(a[:])
	//fmt.Println(sum)
	//testSliceCap()
	//testCopy()
	//array1  := [...]int{1, 2, 3, 4}
	//fmt.Println(array1)
	//testAppend()
	//testStrSlice()
	testStrReverseUtf8()
}

