package main

import (
	"fmt"
	"time"
	"math/rand"
)

// 计算斐波纳契数 1 1 2 3 5 8 13 21 34
func fib(num float64) float64 {
	x, y := 1.0, 1.0
	for i := 0; i < int(num); i++ {
		x, y = y, x+y
	}
	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)
	return x
}

func main() {
	start := time.Now()
	for i := 0; i < 15; i++ {
		n := fib(float64(i))
		fmt.Printf("Fib(%v): %v\n", i, n)
	}

	elapsed := time.Since(start)
    fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}