package main 

import (
	"zuojie.com/day5/pipeline"
	"fmt"
)

func main(){
	p := pipeline.ArraySource(3,2,6,7,4,1)
	// for {
	// 	if num,ok := <-p ;ok{
	// 		fmt.Println(num)
	// 	}else{
	// 		break
	// 	}
	// }
	for v := range p{
		fmt.Println(v)
	}
}