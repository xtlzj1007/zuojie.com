package main

import (
	"fmt"
)

type Phone struct{
	PayMap map[string]Pay
}


func (p *Phone) OpenWeChayPay(){
	weChatPay := &WeChatPay{}
	p.PayMap["wechat_pay"] = weChatPay 
}

func (p *Phone) OpenAilPay(){
	aliPay := &AliPay{}
	p.PayMap["ali_pay"] = aliPay
} 

func (p *Phone) PayMoney(name string,money float32) (err error){
	pay,ok := p.PayMap[name]
	if !ok {
		err = fmt.Errorf("不支持[%s]支付方式",name)
		return 
	}
	pay.pay(1023,money)
	return   
}