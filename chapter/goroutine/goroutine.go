package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func Loop() {
	for i := 1; i < 11; i++ {
		time.Sleep(time.Second * 1)
		fmt.Printf("%d,", i)
	}
}

func Loop2() {
	for i := 1; i < 11; i++ {
		time.Sleep(time.Microsecond * 30)
		fmt.Printf("%d,", i)
	}
}

var chanInt chan int = make(chan int, 10)

func Send() {
	chanInt <- 1
	time.Sleep(time.Second * 1)
	chanInt <- 2
	time.Sleep(time.Second * 1)
	chanInt <- 3
}

func Receive() {
	// num := <-chanInt
	// fmt.Println("num: ", num)
	// num = <-chanInt
	// fmt.Println("num: ", num)
	// num = <-chanInt
	// fmt.Println("num: ", num)

	for {
		select {
		case num := <-chanInt:
			fmt.Println("num: ", num)
		}
	}
}

// 协程同步
var WG sync.WaitGroup

func Read() {
	for i := 1; i < 3; i++ {
		WG.Add(1)
	}
}
func Write() {
	for i := 1; i < 3; i++ {
		fmt.Println("done: ", i)
		time.Sleep(time.Second * 1)
		WG.Done()
	}
}
