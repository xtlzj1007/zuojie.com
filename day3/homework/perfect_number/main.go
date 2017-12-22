package main

import "fmt"

//1000以内的完数
func main() {
	var n int //下标
	var s int //用来计算因子之和
	for i := 2; i < 1000; i++ {
		k := [100]int{}
		n, s = -1, 0 //每次计算时，初始化n,s
		for j := 1; j < i; j++ {
			if i%j == 0 { //获取因子
				n ++
				s += j //计算因子之和
				k[n] = j
			}
		}
		if s == i {
			fmt.Printf("%d=", i) //输出完数
			for i := 0; i < n; i++ {
				fmt.Printf("%d+", k[i]) //输出因子
			}
			fmt.Println(k[n]) //输出最后一个因子并换行
		}
	}
}
