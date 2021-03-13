package main

import "fmt"

// 指针数组
func TestPointArr() {
	a, b := 1, 2
	pointArr := [...]*int{&a, &b}
	fmt.Println("指针数组：", pointArr)
}

// 数组指针
func TestArrPoint() {
	arr := [...]int{3, 4, 5}
	arrPoint := &arr
	fmt.Println("数组指针：", arrPoint)
}

func test() {
	fmt.Print("1111111")
}
