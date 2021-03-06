package main

import (
	"fmt"
	"encoding/json"
)

type Student struct{
	Name string `json:"name"` //在json中的字段名称
	Age int     `json:"age"`
	Sex string  `json:"sex"`
}

func main(){
	var s Student
	s.Age = 200
	s.Name = "abc"
	s.Sex = "man"
	data ,err := json.Marshal(s) //创建一个json字符串
	if err != nil{
		fmt.Printf("json marshal failed,err:%v\n",err)
		return
	}
	fmt.Printf("%s\n",data)
	var s1 Student
	err = json.Unmarshal(data,&s1) //反序列化
	if err != nil{
		fmt.Printf("json Unmarshal failed,err:%v",err)
		return
	}
	fmt.Printf("s1:%#v\n",s1)

}