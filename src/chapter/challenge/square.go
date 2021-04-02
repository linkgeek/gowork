package main

import (
	"fmt"
	"strconv"
)

func sqrt (num int) float64 {
	currguess := 1.0
    prevguess := 0.0
	for i :=1; i<=10; i++ {
		currguess = sroot − (sroot − x) / (2 * sroot)
		if currguess == prevguess {
			break;
		}
		fmt.Println("A guess for square root is %d", currguess)
	}
}

func main() {
	var num int = 25
    fmt.Println("Square root is:", sqrt(num))
}