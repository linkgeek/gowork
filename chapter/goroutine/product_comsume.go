package main

import (
	"fmt"
	"time"
)

var (
	infos = make(chan int, 10)
)

func Producer(idx int) {
	infos <- idx
}

func Consumer(idx int) {
	fmt.Printf("Consumer: %d, Receive: %d\n", idx, <-infos)
}

func main() {
	// 正常
	for i := 0; i < 10; i++ {
		go Producer(i)
	}
	for i := 0; i < 10; i++ {
		go Consumer(i)
	}

	// 生产者 > 消费者
	// for i := 0; i < 100; i++ {
	// 	go Producer(i)
	// }
	// for i := 0; i < 10; i++ {
	// 	go Consumer(i)
	// }

	// 消费者 > 生产者
	// for i := 0; i < 10; i++ {
	// 	go Producer(i)
	// }
	// for i := 0; i < 100; i++ {
	// 	go Consumer(i)
	// }

	time.Sleep(time.Second * 5)
}
