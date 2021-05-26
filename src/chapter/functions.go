package main

import (
    "os"
    "strconv"
)

// main 函数
func main() {
    number1, _ := strconv.Atoi(os.Args[1])
    number2, _ := strconv.Atoi(os.Args[2])
    println("Sum:", number1+number2)

	sum, mul := calc(os.Args[1], os.Args[2])
    println("Sum:", sum)
    println("Mul:", mul)

	firstName := "John"
    updateName(firstName)
    println(firstName)
	updateName2(&firstName)
    println(firstName)
}

// 自定义函数
func sum(number1 string, number2 string) int {
    int1, _ := strconv.Atoi(number1)
    int2, _ := strconv.Atoi(number2)
    return int1 + int2
}
func sum2(number1 string, number2 string) (result int) {
    int1, _ := strconv.Atoi(number1)
    int2, _ := strconv.Atoi(number2)
    result = int1 + int2
    return
}

// 返回多个值
func calc(number1 string, number2 string) (sum int, mul int) {
    int1, _ := strconv.Atoi(number1)
    int2, _ := strconv.Atoi(number2)
    sum = int1 + int2
    mul = int1 * int2
    return
}

// 更改函数参数值（值传递）
func updateName(name string) {
    name = "David"
}

// 更改函数参数值（指针:内存地址）
func updateName2(name *string) {
    *name = "Piter"
}