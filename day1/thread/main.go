package main
import ("fmt"
		"time")


func PrintEven(n int){
	for i := 0; i < n; i++{
		if i % 2 == 0 {
			fmt.Printf("%d\n",i)
		}
	}
}

func PrintOdd(n int){
	for i := 0; i < n; i++{
		if i % 2 != 0 {
			fmt.Printf("%d\n",i)
		}
	}
}


func main(){
	go PrintEven(10) // 启动偶数子线程
	go PrintOdd(10) // 启动奇数子线程
	
	sum,sub := calc(21,31)
	fmt.Println(sum,sub)
	time.Sleep(10 * time.Second) // 1是1纳秒
}

func calc(a int, b int)(int,int){
	sum := a+b
	sub := a-b
	return sum,sub
}



