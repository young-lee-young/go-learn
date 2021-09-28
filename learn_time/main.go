package main

import (
	"time"
	"fmt"
)

func baseTime() {
	// 1纳秒，十亿分之一秒
	nanoSecond := time.Duration(1)
	fmt.Println(nanoSecond)		// 1ns
	fmt.Println(nanoSecond == 1) 	 // true
	fmt.Println(nanoSecond * 1000000000 == time.Second)		// true
}

func main() {
	baseTime()
}
