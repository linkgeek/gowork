package main

import "fmt"

func givemeanumber() int {
    return -1
}

func main() {
	x := 27
    if x%2 == 0 {
        fmt.Println(x, "is even")
    }
	
    if num := givemeanumber(); num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has only one digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

	//fmt.Println(num)
}