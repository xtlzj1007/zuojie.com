package main

import "fmt"

var is_login bool

type add_func func(int, int) int

func add(a, b int) int {
	return a + b
}

func operator(op add_func, a int, b int) int {
	return op(a, b)
}

type handle func()

func login(run handle) {
	if is_login == false {
		return
	}
}

func TestOrder() {

}

func TestAccount() {

}

func TestLogin() {
	login(TestAccount)
	login(TestOrder)
}


func main() {
	c := add
	//fmt.Printf("c  :%p type(c)  :%T\nadd:%p type(add):%T\n", c, c, add, add)
	//sum := c(10, 20)
	sum := operator(c, 10, 20) //函数式编程
	fmt.Println(sum)
	//sum = add(10, 20)
	//fmt.Println(sum)
}
