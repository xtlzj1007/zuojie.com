package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	modify(arr1)
	//arrLength := len(arr1)
	//fmt.Println(arrLength)
	fmt.Println("In main(),array values:", arr1)
	//for i, v := range arr1 { //从数组中取出数据
	//	fmt.Println(i, v)
	//}
}

func modify(array [5]int) {
	array[0] = 10
	fmt.Println("In modify(),array values:", array)
}
