package main

import "fmt"

func testMap() {
	var a map[string]int
	a = make(map[string]int, 100)
	a["123"] = 200
	fmt.Println(a)
	for k, v := range a {
		fmt.Printf("a[%s] = %d\n", k, v)
	}
	val, ok := a["123"]
	if ok {
		fmt.Printf("val = %d\n", val)
	} else {
		fmt.Println("not found")
	}
	val = a["aaa"]
	fmt.Println(val)
	modify(a)
	fmt.Println(a)
	items := make([]map[int]int, 5)
	for i := 0; i < 5; i++ {
		items[i] = make(map[int]int)
	}

}

func modify(a map[string]int) {
	a["one"] = 123
}

func main() {
	testMap()
}
