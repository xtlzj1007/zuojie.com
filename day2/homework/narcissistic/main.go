package main

import "math"

func main() {
	PrintNarcissistic()
}

//计算水仙花数
func PrintNarcissistic() {
	for i := 100; i < 1000; i++ {
		hundreds := i / 100   // 123 / 100 = 1
		decade := i / 10 % 10 // 123 / 10 = 12 % 10 = 2
		unit := i % 10        // 123 % 10 = 3
		if math.Pow(float64(hundreds), 3)+math.Pow(float64(decade), 3)+math.Pow(float64(unit), 3) == float64(i) {
			println(i)
		}
	}
}

