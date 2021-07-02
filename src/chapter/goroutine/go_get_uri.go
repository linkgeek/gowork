package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    apis := []string{
        "https://management.azure.com",
        "https://dev.azure.com",
        "https://api.github.com",
        "https://outlook.office.com/",
        "https://api.somewhereintheinternet.com/",
        "https://graph.microsoft.com",
    }

    //testBasic(apis)
    testChanApi(apis)
    //testBufferChan2()
}

func checkApi(api string) {
    _, err := http.Get(api)
    if err != nil {
        fmt.Printf("ERROR: %s is down!\n", api)
        return
    }

    fmt.Printf("SUCCESS: %s is up and running!\n", api)
}
func testBasic(apis []string) {
    start := time.Now()
    for _, api := range apis {
        go checkApi(api)
    }

    elapsed := time.Since(start)
    fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())

    time.Sleep(2 * time.Second)
}


// 无缓冲 channel
func checkChanApi(api string, ch chan string) {
    _, err := http.Get(api)
    if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
        return
    }

    ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}
func testChanApi(apis []string) {
    start := time.Now()
	ch := make(chan string)
	//ch := make(chan string, 10)
    for _, api := range apis {
        go checkChanApi(api, ch)
    }

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

    elapsed := time.Since(start)
    fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

// 有缓冲 channel
func send(ch chan string, msg string) {
    ch <- msg
}
func testBufferChan() {
    size := 4
    ch := make(chan string, size)
    send(ch, "one")
    send(ch, "two")
    send(ch, "three")
    send(ch, "four")
    fmt.Println("All data sent to the channel ...")
    for i := 0; i < size; i++ {
        fmt.Println(<-ch)
    }
    fmt.Println("Done!")
}
func testBufferChan2() {
    size := 2
    ch := make(chan string, size)
    send(ch, "one")
    send(ch, "two")
    go send(ch, "three")
    go send(ch, "four")
    fmt.Println("All data sent to the channel ...")
    for i := 0; i < 4; i++ {
        fmt.Println(<-ch)
    }
    fmt.Println("Done!")
}