package main

import (
	"fmt"
)


func main(){
	phone := &Phone{
		PayMap:make(map[string]Pay,16),
	}
	phone.OpenAilPay()
	err := phone.PayMoney("ali_py",200)
	if err != nil{
		fmt.Printf("支付失败,失败原因%v",err)
		fmt.Println("使用支付宝支付")
		err = phone.PayMoney("ali_py",20.22)
		if err != nil{
			fmt.Printf("支付失败，失败原因:%v\n",err)
			return
		}
	}
	fmt.Println("支付成功，欢迎下次光临！")
}
