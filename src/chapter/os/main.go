package main

import (
	"os"
    "strconv"
	"fmt"
)

func getSum(num1 string, num2 string) int {
	int1, _ := strconv.Atoi(num1)
    int2, _ := strconv.Atoi(num2)
    return int1 + int2
}

func calc(num1 string, num2 string) (sun int, mul int) {
	int1, _ := strconv.Atoi(num1)
    int2, _ := strconv.Atoi(num2)
    sum = int1 + int2
    mul = int1 * int2
	return
}

func test() {
	fmt.Print(os.Args)
	sum := getSum(os.Args[1], os.Args[2])
    println("Sum:", sum)

	sum, mul := calc(os.Args[1], os.Args[2])
	println("Sum:", sum)
	println("Mul:", mul)
}

func main() {
	test()
}
