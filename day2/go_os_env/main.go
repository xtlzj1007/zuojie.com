package main

import (
	"os"
	"fmt"
)

func main() {
	path, _ := os.LookupEnv("path")
	hostname,_ := os.Hostname()
	fmt.Printf("%v\n%v\n", path, hostname)
	fmt.Printf("abc"+"def")
}
