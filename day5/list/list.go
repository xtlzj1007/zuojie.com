package main

import (
	"fmt"
)

type Teacher struct{
	Name string
	Age int
	Next *Teacher
}


func createList(){
    var header * Teacher = &Teacher{}
	header.Age = 200
	header.Name = "a"
	printList(header)
	p := new(Teacher)
	p.Name = "b"
	p.Age = 100

	header.Next = p 
	printList(header)
}

func createInHeader(h *Teacher,name string,age int)(*Teacher){
	//链表从头部插入数据
	p := &Teacher{}
	p.Age = age
	p.Name = name 
	p.Next = h 
	return p 
}

func createInTail(tail *Teacher,name string,age int)(*Teacher){
	p := &Teacher{}
	p.Age = age
	p.Name = name 
	if tail == nil{
		return p 
	}else{
    	tail.Next = p
		return p
	}
}

func printList(h *Teacher){
	for h!= nil{
		fmt.Printf("Name:%v Age:%v\n",h.Name,h.Age)
		h = h.Next
	}
}

func TestCreateInHeader(){
	var header * Teacher
	header = createInHeader(header,"a",12)
	header = createInHeader(header,"b",20)
	header = createInHeader(header,"c",33)
	printList(header)
}

func TestCreateInTail(){
	// var header * Teacher
	tail := &Teacher{}
	tail = createInTail(tail,"a",10)
	tail = createInTail(tail,"b",12)
	tail = createInTail(tail,"c",13)
	printList(tail)
}

//构造函数
func NewTeacher(name string,age int) *Teacher{
	// return &Teacher{
	// 	Name:name,
	// 	Age:age,
	// }
	p := new(Teacher)
	p.Name = name
	p.Age = age
	return p 
}

func main(){
	// createList()
	// TestCreateInHeader()
	TestCreateInTail()
}

