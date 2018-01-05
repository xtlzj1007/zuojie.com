package pipeline

import "sort"

// 阵列源 只能拿东西
func ArraySource(a ...int) <-chan int{
	out := make(chan int)
	go func(){
		for _,v := range a {
			out <- v
		}
		close(out)
	}()
	return out 
}

func InMemSort(int <-chan int) <-chan int{
	out := make(chan int)
	go func(){
		a := [] int{}
		for v := range in {
			a = append(a,v)
		} 

		sort.Ints(a)
		for _,v:= range a{
			out <- v
		}
		close(out)
	}()
	return out
}
