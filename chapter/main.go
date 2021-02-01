package main

import (
	"fmt"
	"impala/chapter/learn1"
	"impala/chapter/show2"
)

func init() {
	fmt.Print("在main函数执行之前\n")
}

func main() {
	learn1.Learn1()
	show2.Show2()
	fmt.Print("this is main")
}
