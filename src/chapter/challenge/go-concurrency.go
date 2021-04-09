package main

import (
	"fmt"
	"time"
	"math/rand"
)

// 计算斐波纳契数 1 1 2 3 5 8 13 21 34
func fib(num float64, ch chan string) {
	x, y := 1.0, 1.0
	for i := 0; i < int(num); i++ {
		x, y = y, x+y
	}
	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	ch <- fmt.Sprintf("Fib(%v): %v\n", num, x)
}

func main() {
	start := time.Now()
	
	size := 15
	ch := make(chan string, size)
	for i := 0; i < size; i++ {
		go fib(float64(i), ch)
	}

	for i := 0; i < size; i++ {
        fmt.Printf(<-ch)
    }

	elapsed := time.Since(start)
    fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}