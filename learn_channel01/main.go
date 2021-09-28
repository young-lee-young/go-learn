package main

import (
	"fmt"
	"time"
)

/**
	没有缓冲区的channel，收数据或者发数据会发生阻塞
 */

func numFunc01(ch chan int) {
	fmt.Println("没有人发数据，阻塞中...")
	num := <- ch
	fmt.Println("收取数据：", num)
}

func numFunc02(ch chan int) {
	fmt.Println("没有人收数据，阻塞中...")
	ch <- 2
	fmt.Println("有人收数据了，阻塞结束")
}

func main() {
	ch := make(chan int)

	// 测试发数据阻塞
	go numFunc01(ch)
	time.Sleep(time.Second * 3)
	fmt.Println("发数据啦")
	time.Sleep(time.Second * 1)
	ch <- 1
	time.Sleep(time.Second * 1)

	// 测试收数据阻塞
	go numFunc02(ch)
	time.Sleep(time.Second * 3)
	fmt.Println("收数据啦")
	time.Sleep(time.Second * 1)
	num := <- ch
	fmt.Println("收取数据：", num)
	time.Sleep(time.Second * 1)
}
