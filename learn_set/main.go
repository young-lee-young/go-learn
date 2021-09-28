package main

import (
	"fmt"
	"learn/learn_set/set"
)

/**
	Go实现set
 */

func main() {
	setNew := set.New(1, 2)
	setNew.Add("lee")
	setNew.Add("zhao")
	fmt.Println(setNew.Size())
	fmt.Println(setNew.Contains("lee"))
	setNew.Remove("lee")
	fmt.Println(setNew.Size())
	fmt.Println(setNew.Contains("lee"))
	setNew.Clear()
	fmt.Println(setNew.Size())
}
