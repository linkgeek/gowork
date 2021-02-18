package main

import "fmt"

func TestChan() {
	// 初始化channel
	// channel在使用之前，必须进行make初始化
	// 否则，它会是一个nil
	ch := make(chan int)

	// fmt.Println(ch)

	// 输出channel
	go func() {
		fmt.Println(<-ch)
	}()

	// 输入channel
	ch <- 1

	// 关闭channel
	close(ch)

	fmt.Println(ch)
	// ch <- 2

	<-ch
}
