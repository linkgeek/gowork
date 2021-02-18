package main

import (
	"fmt"
	"sync"
)

// 100个人抢购10个鸡蛋
func main0() {
	// 初始化eggs
	eggs := make(chan int, 10)

	// 输入10个鸡蛋
	for i := 0; i < 10; i++ {
		eggs <- i
	}

	// 初始化wg
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int) {
			select {
			case egg := <-eggs:
				fmt.Printf("People : %d, Get egg : %d\n", num, egg)
			default:
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
