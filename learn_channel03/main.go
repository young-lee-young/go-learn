package main

import (
	"fmt"
	"time"
)

/**
	channel关闭后，继续向channel里发送数据，会panic

	任何情况下都不应该在接收端关闭channel
 */

func receiver(ch chan int) {
	for {
		num, ok := <- ch
		fmt.Println("receive num", num, ok)
	}
}

func main() {
	ch := make(chan int)
	go receiver(ch)

	nums := []int{1, 2, 3}
	for _, item := range nums {
		fmt.Println("send num", item)
		ch <- item
		time.Sleep(time.Second * 1)
		// 向关闭的channel发送数据会panic
		// panic: send on closed channel
		if item == 2 {
			fmt.Println("channel closed")
			close(ch)
		}
	}
	time.Sleep(time.Second * 5)
}
