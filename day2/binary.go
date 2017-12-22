package main

import (
	//h "zuojie.com/day1/homework"
	f "zuojie.com/day2/fibo"
	"fmt"
)

func main()  {
	//h.PrintData()
	f.PrintFibo(10)
	fmt.Println()
	f.PrintFibo(15)
}

//type User struct {
//	name string
//	flag uint8
//}
//
//func is_set_flag(user User, isSet bool, flag uint8) User {
//	if isSet == true {
//		user.flag = user.flag | flag
//	} else {
//		user.flag = user.flag ^ flag
//	}
//	return user
//}
