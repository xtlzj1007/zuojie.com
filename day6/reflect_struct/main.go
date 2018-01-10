package main

import (
	"fmt"
	"reflect"
)


type Student struct{
	Name string
	Age int 
	Sex int 
}


func (s *Student) Set(name string,age int ,sex int){
	s.Name = name
	s.Age = age
	s.Sex = sex
}

func (s *Student) GetName(name string) {
	s.Name = name 
}

func testStruct(){
	var stu *Student = &Student{}
	stu.Set("jim",18,1)
	valueinfo := reflect.ValueOf(stu)

	fieldNum := valueinfo.Elem().NumField()
	fmt.Println("field name:",fieldNum)
	// sexValueInfo := valueinfo.Elem().FieldByName("Name")
	// fmt.Println("sex=",sexValueInfo.Int())
	// sexValueInfo.SetInt(100)
	// fmt.Println(stu)

	// setMethod := valueinfo.MethodByName("Set")
	// fmt.Println(setMethod)

	// var params []reflect.Value
	// name := "Tom"
	// age := 1000
	// sex := 2332

	// params = append(params,reflect.ValueOf(name))
	// params = append(params,reflect.ValueOf(age))
	// params = append(params,reflect.ValueOf(sex))

	// setMethod.Call(params)
	// fmt.Println(stu)

}

func main(){
	testStruct()
}