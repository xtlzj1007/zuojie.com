package main

import (
	"log"
	"os"
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

func Json_demo(){
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

func Json_demo2(){
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
		if k != "Title" {
			v[k] = false
			}
		}
		if err := enc.Encode(&v); err != nil {	
		log.Println(err)
		}
	}
}

func main(){
	// Json_demo()
	Json_demo2()
}
