package main

import (
	"github.com/myuser/calculator"
	"rsc.io/quote"
	"fmt"
	"time"
)

func init() {
	x := 27
    if x%2 == 1 {
        fmt.Println(x, "is even")
    }
	fmt.Println(time.Now().Weekday().String())
}

func main() {
    total := calculator.Sum(3, 5)
    println(total)
    println("Version: ", calculator.Version)

	println(quote.Hello())

	// total := calculator.internalSum(5)
    // println(total)
    // println("Version: ", calculator.logMessage)
}