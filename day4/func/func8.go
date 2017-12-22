package main

import (
	"strings"
	"fmt"
)

func makeSufifixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	func1 := makeSufifixFunc(".bmp")
	func2 := makeSufifixFunc(".jpg")
	fmt.Println(func1("test"))
	fmt.Println(func2("test"))
}
