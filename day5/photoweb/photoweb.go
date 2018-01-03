package main 

import (
	"fmt"
	"log"
	"net/http"
)

const (
	UPLOAD_DIR = "./uploads"
)

//支持图片上传
func uploadHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		fmt.Fprint(w, `<div><form method="POST" 
			action="\upload\" enctype="multipart/form-data">
			Choose an image to upload: <input name="image" type="file" />
			<input type="submit" value="Upload" /></form></div>`)
		// fmt.Fprintln(w,"<h1>hello world</h1>")
		return 
	}
	if r.Method == "POST"{
		
	}
	
}

//在网页里可以上传图片
//能看到网页上的图片列表
//可以删除指定的图片

func main(){
	http.HandleFunc("/upload",uploadHandler)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal("ListenAndServer:",err.Error())
	}
}