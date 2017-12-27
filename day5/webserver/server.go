package main

import (
	"fmt"
	"net/http"
)

//一个函数web输出函数
func HelloWorld(writer http.ResponseWriter,request *http.Request){
	fmt.Fprintf(writer,"<h1>Hello World  %s!</h1>",request.FormValue("name"))
}

func main(){
	http.HandleFunc("/",HelloWorld) //路由，并把函数结果发送出去
	http.ListenAndServe(":8888",nil) //监听地址
}