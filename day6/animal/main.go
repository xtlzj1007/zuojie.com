package main

// 接口定义

import (
	"fmt"
)

// 动物接口
type Animal interface {
	Eat()
	Talk()
}

// 定义dog类
type Dog struct {
	Name string
}

func (d *Dog) Eat(){
	fmt.Println(d.Name,"eating")
}

func (d *Dog) Talk(){
	fmt.Println(d.Name,"wangwang")
}

type Cat struct{
	Name string
}

func (c *Cat) Eat(){
	fmt.Println(c.Name,"cat eating")
}

func (c *Cat) Talk(){
	fmt.Println(c.Name,"cat miaomiao")
}


func Test1(){
	var a Animal
	var d Dog
	var c Cat
	// d.Eat()
	a = &d   // 将dog对象存入animal接口
	a.Eat()  // 调用接口中的方法
	a.Talk()
	// fmt.Println()
	a = &c
	a.Eat()
	a.Talk() 
}

func TestOperator(){
	var animalList []Animal
	d1 := &Dog{
		Name:"小黑",
	}
	animalList = append(animalList,d1)

	d2 := &Dog{
		Name:"布丁",
	}
	animalList = append(animalList,d2)

	c1 := &Cat{
		Name:"ruby",
	}
	animalList = append(animalList,c1)

	for _,v := range animalList{
		v.Eat()
		v.Talk()
	}
}


func main(){
	TestOperator()
}