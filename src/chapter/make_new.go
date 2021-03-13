package main

import (
	"fmt"
	"reflect"
)

// 创建切片
func makeSlice() {
	mSlice := make([]string, 3)
	mSlice[0] = "dog"
	mSlice[1] = "cat"
	mSlice[2] = "pig"
	//mSlice[3] = "pig3"
	fmt.Print(mSlice)
}

// 创建map, 返回的是引用
func makeMap() {
	mMap := make(map[int]string)
	mMap[10] = "dog"
	mMap[100] = "dog2"
	fmt.Print(mMap, "\n")
	fmt.Print(reflect.TypeOf(mMap), "\n")
}

// 创建map, 返回的是指针
func newMap() {
	newMap := new(map[int]string)
	fmt.Print(reflect.TypeOf(newMap), "\n")
}

// 创建chan
func makeChan() {
	mChan := make(chan int)
	close(mChan)
}

func main11() {
	//makeSlice()
	makeMap()
	newMap()
}
