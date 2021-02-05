package main

import "fmt"

// 向切片中添加元素
func appendEleSlice() {
	mSlice := make([]string, 2)
	mSlice[0] = "dog"
	mSlice[1] = "cat"

	fmt.Print("len=", len(mSlice), "\n")
	fmt.Print("cap=", cap(mSlice), "\n")

	mSlice = append(mSlice, "pig")
	fmt.Print(mSlice, "\n")
	fmt.Print("After len=", len(mSlice), "\n")
	fmt.Print("After cap=", cap(mSlice), "\n")
}

// 拷贝切片元素
func copySlice() {
	mSlice1 := make([]string, 2)
	mSlice1[0] = "id-1-1"
	mSlice1[1] = "id-1-2"

	mSlice2 := make([]string, 3)
	mSlice2[0] = "id-2-1"
	mSlice2[1] = "id-2-2"
	mSlice2[2] = "id-2-3"

	copy(mSlice1, mSlice2)
	fmt.Print(mSlice1, "\n")
}

// 删除
func delFromMap() {
	mMap := make(map[int]string)
	mMap[10] = "dog"
	mMap[100] = "dog2"
	delete(mMap, 10)
	fmt.Print(mMap, "\n")
}

func main12() {
	//appendEleSlice()
	//copySlice()
	delFromMap()
}
