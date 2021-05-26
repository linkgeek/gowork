package main

// 声明变量
var firstName string
var firstName, lastName string
var age int
var (
    firstName, lastName string
    age int
)
// 初始化变量
var (
    firstName string = "John"
    lastName  string = "Doe"
    age       int    = 32
	featureFlag bool = true
)
var (
    firstName = "John"
    lastName  = "Doe"
    age       = 32
	enable    = false
)
// 通过多种方式初始化变量
var (
    firstName, lastName, age = "John", "Doe", 32
)


// 声明常量
const HTTPStatusOK = 200
const (
    StatusOK              = 0
    StatusConnectionReset = 1
    StatusOtherError      = 2
)

func main() {
    firstName, lastName := "John", "Doe"
    age := 32
    println(firstName, lastName, age)
}