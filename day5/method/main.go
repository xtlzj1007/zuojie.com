package main

import (
	"fmt"
)

type Int int

type Student struct{
	Name string `json:"name"` //在json中的字段名称
	Age int     `json:"age"`
	Sex string  `json:"sex"`
}

// func Add(a,b int) int{   //函数
// 	return a + b
// }

func (i *Int)Add(a,b int) { //方法
	// return a + b
	*i = Int(a+b) //给对象本身赋值
	return  
}


func TestInt(){
	var a Int
    a.Add(100,200)	
	fmt.Println(a)
}

func (s *Student) Set(name string,age int){ //给类型添加方法
	s.Name = name
	s.Age = age 
}

func main(){
	var s Student
	s.Set("zj",23)
	fmt.Println(s)	
}


