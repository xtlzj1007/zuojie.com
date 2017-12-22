package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	PrintTenRanNum()
}

// 输出随机数
func PrintTenRanNum() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		fmt.Println(rnd.Intn(100), rnd.Float32())
	}
}
