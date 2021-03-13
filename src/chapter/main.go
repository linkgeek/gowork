package main

import (
	"fmt"
	"runtime"
	"chapter/http"
)

func init() {
	fmt.Print("在main函数执行之前\n")
}

func main() {
	//learn1.Learn1()
	//show2.Show2()
	fmt.Println("this is main")

	fmt.Println("cpu核数=", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	//go goroutine.Loop()
	//go goroutine.Loop2()

	// 协程通信
	//go goroutine.Send()
	//go goroutine.Receive()

	// 协程同步
	// goroutine.Read()
	// go goroutine.Write()
	// goroutine.WG.Wait()
	// fmt.Println("All Done!")

	test.TestHttp()

	//time.Sleep(time.Second * 30)
}
