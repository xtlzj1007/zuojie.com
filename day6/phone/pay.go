package main


// 支付的接口
type Pay interface{
	pay(user_id int64,money float32)  (err error)
}

