package main

import (
	"fmt"
	"time"
)

/**
	channel关闭后继续接收数据，不会报错
 */

func receiver(ch chan int) {
	time.Sleep(time.Second * 1)
	for {
		// 从关闭的channel中接收数据
		// channel中没有数据后，ok会返回false
		num, ok := <- ch
		fmt.Println("receive num", num, ok)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	ch := make(chan int, 3)
	go receiver(ch)

	nums := []int{1, 2, 3}
	for _, item := range nums {
		fmt.Println("send num", item)
		ch <- item
	}
	close(ch)
	fmt.Println("channel closed")
	time.Sleep(time.Second * 5)
}
