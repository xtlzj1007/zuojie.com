package main

import (
	"fmt"
)

type People struct{
	Name string
	Age int
	int
}

type Student struct{
	Score int
	People // 使用继承
}


func main(){
	var s Student
	s.Name = "abc"
	s.Age = 100
	s.int = 1024
	s.Score = 99
	fmt.Printf("%#v\n",s)
}


