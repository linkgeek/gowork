package main

import (
    "fmt"
    "moduledemo/mypackage"  // 导入同一项目下的mypackage包
	"mypackage2"
)
func main() {
    mypackage.New()
	mypackage2.New()
    fmt.Println("main")
}