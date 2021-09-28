package main

import (
	"fmt"
)

/**
	关闭已经关闭的channel，会panic
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
		if item == 2 || item == 3 {
			// 多次关闭channel
			// close of closed channel
			close(ch)
			fmt.Println("channel closed")
		} else {
			fmt.Println("send num", item)
			ch <- item
		}
	}
}
