package homework

import "fmt"

func PrintData(){
	str       := "Hello World!"
	num       := 1024
	float_num := 3.14
	fmt.Printf("字符串：%s\n1024整数：%d\n1024二进制：%b\n1024十六进制：%x\n1024八进制：%o\n浮点数:%.2f\n",str,num,num,num,num,float_num)
}