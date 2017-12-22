package main
import "zuojie.com/day1/homework"
import "time"

func main(){
	go homework.PrintData()
	time.Sleep(3 * time.Second)
}