package main

import (
	"fmt"
)

type Student struct{
	Name string
	Age int
	int
}


func main(){
	var s Student
	s.Name = "abc"
	s.Age = 100
	s.int = 1024
	fmt.Printf("%#v\n",s)
}