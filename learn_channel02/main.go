package main

import (
	"fmt"
	"time"
)

/**
	带缓冲区的channel
 */

func numFunc01(ch chan int) {
	//time.Sleep(time.Second * 1)
	num := <- ch
	fmt.Println("收取数据：", num)
}

func main() {
	//ch := make(chan int, 1)
	ch := make(chan int, 2)

	go numFunc01(ch)

	ch <- 1
	// 会立即发送数据1，不阻塞
	fmt.Println("发送数据 1")
	time.Sleep(time.Second * 1)
	// 缓冲区大小为1，这里阻塞，知道缓冲区再次空出位置
	ch <- 2
	fmt.Println("发送数据 2")
	time.Sleep(time.Second * 2)
}
