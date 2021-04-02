package main

import (
	"fmt"
	"os"
    "io"
)

func test() {
	for i := 1; i <= 4; i++ {
        defer fmt.Println("deferred", -i)
        fmt.Println("regular", i)
    }
}

func typicalCase() {
	f, err := os.Create("notes.txt")
    if err != nil {
        return
    }
    defer f.Close()

    if _, err = io.WriteString(f, "Learning Go!"); err != nil {
        return
    }

    f.Sync()
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic("Panic in g() (major)")
    }
    defer fmt.Println("Defer in g()", i)
    fmt.Println("Printing in g()", i)
    g(i + 1)
}

func main() {
    //g(0)
    //fmt.Println("Program finished successfully!")

    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in main", r)
        }
    }()
    g(0)
    fmt.Println("Program finished successfully!")
}