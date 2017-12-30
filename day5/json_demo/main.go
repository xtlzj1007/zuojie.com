package main

import (
	"fmt"
	"encoding/json"
)

type Book struct{
	Title string     `json:"title"`
	Authors []string `json:"authors"`
	Publisher string `json:"publisher"`
	IsPublished bool `json:"is_published"`
	Price float32    `json:"price"`
}

func main(){
	var gobook Book
	gobook.Title = "Go语言编程"
	gobook.Authors = []string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo"}
	gobook.Publisher = "ituring.com.cn"
	gobook.IsPublished = true
	gobook.Price =  9.99
	// 序列化数据成json字符串
	b,err := json.Marshal(gobook)
	if err == nil {
		fmt.Printf("%s\n",b)
	}else{
		fmt.Println(err)
	}
	var r Book
	// 反序列化数据成go对象
	err = json.Unmarshal(b,&r)
	if err == nil{
		fmt.Println(r)
	}
}
