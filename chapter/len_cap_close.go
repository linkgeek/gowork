package main

import "fmt"

// len -> string、array、slice、map、chan
// cap -> slice、array、chan
// close -> chan

// 长度、容量
func getLenCap() {
	mSlice := make([]string, 2)
	mSlice[0] = "id-1-1"
	mSlice[1] = "id-1-2"
	fmt.Println("len=", len(mSlice))
	fmt.Println("cap=", cap(mSlice))

	mSlice = append(mSlice, "id-1-3")
	mSlice = append(mSlice, "id-1-4")
	mSlice = append(mSlice, "id-1-5")

	fmt.Println("len=", len(mSlice))
	fmt.Println("cap=", cap(mSlice))
}

// 创建、关闭chan
func createCloseChan() {
	mChan := make(chan int, 1)
	mChan <- 1
	defer close(mChan)
	//mChan <- 2
}

func main13() {
	//getLenCap()
	createCloseChan()
}
