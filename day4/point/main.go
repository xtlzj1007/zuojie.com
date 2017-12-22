package main

import "fmt"

func modify_arr1(arr *[6]int) {
	fmt.Printf("%p\n", arr)
	arr[0] = 100
	fmt.Println(*arr)
}

func modify_arr2(arr []int) {
	arr[0] = 100
}

func main() {
	arr1 := [6]int{}
	modify_arr1(&arr1)
	fmt.Printf("%p\n", &arr1)
	fmt.Println(arr1)

}
