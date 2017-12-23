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
	p := &Teacher{}
	p.Age = age
	p.Name = name 
	p.Next = h 
	return p 
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
	printList(header)
}

func main(){
	// createList()
	TestCreateInHeader()
}

