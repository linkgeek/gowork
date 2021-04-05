package main2

import (
	"fmt"
	"strconv"
)

// 编写 FizzBuzz 程序
func fizzbuzz(num int) string {
	switch {
	case num%15 == 0:
		return "FizzBuzz"
	case num%5 == 0:
		return "Buzz"
	case num%3 == 0:
		return "Fizz"
	}
	return strconv.Itoa(num)
}

func testSwitch() {
	for num := 1; num <= 100; num++ {
		fmt.Printf("%d - %s\n", num, fizzbuzz(num))
	}
}

// 推测平方根
func sqrt (num float64) float64 {
	prevguess := 0.0
	currguess := 1.0
	for i :=1; i<=10; i++ {
		prevguess = currguess
		currguess = prevguess - (prevguess*prevguess-num)/(2*prevguess)
		if currguess == prevguess {
			break
		}
		fmt.Println("A guess for square root is %d", currguess)
	}
	return currguess
}

func testSqrt() {
	var num float64 = 25
    fmt.Println("Square root is:", sqrt(num))
}

// 要求用户输入一个数字，如果该数字为负数，则进入紧急状态
func numValid()  {
	val := 0
	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)

		switch {
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		case val < 0:
			panic("You entered a negative number!")
		default:
			fmt.Println("You entered: ", val)
		}
	}
}

func main2() {
	numValid()
}