package main

import (
	"fmt"
	"math"
	"os"
	"io"
	"net/http"
	"encoding/json"
	"log"
)

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Square struct {
	size float64
}

// 正方形面积
func (s Square) Area() float64 {
	return s.size * s.size
}

// 正方形周长
func (s Square) Perimeter() float64 {
	return s.size * 4
}

func testSqure() {
	var s Shape = Square{3}
    fmt.Printf("%T\n", s)
    fmt.Println("Area: ", s.Area())
    fmt.Println("Perimeter:", s.Perimeter())
}

type Circle struct {
	radius float64
}

// 圆面积 圆周长
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
// 圆面积 圆周长
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printInfo(s Shape) {
	fmt.Printf("%T\n", s)
    fmt.Println("Area: ", s.Area())
    fmt.Println("Perimeter:", s.Perimeter())
    fmt.Println()
}

// 实现字符串接口
type Person struct {
    Name, Country string
}
func (p Person) String() string {
    return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}
func printStr() {
	rs := Person{"John Doe", "USA"}
    ab := Person{"Mark Collins", "United Kingdom"}
    fmt.Printf("%s\n%s\n", rs, ab)
}

// 自定义writer
type customWriter struct {}
type GithubResp []struct {
	FullName string `json:"full_name"`
}
func (cw customWriter) Write(p []byte) (n int, err error) {
	var resp GithubResp
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil
}

func customResp() {
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    //io.Copy(os.Stdout, resp.Body)
	writer := customWriter{}
	io.Copy(writer, resp.Body)
}


// 编写自定义服务器 API
type dollars float32
func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}
type database map[string]dollars
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func testHttp() {
	db := database{"Go T-Shirt": 25, "Go Jacket": 55}
    log.Fatal(http.ListenAndServe("localhost:8000", db))
}

func main() {
	// var s Shape = Square{3}
	// printInfo(s)

    // c := Circle{6}
	// printInfo(c)

	// 调用：curl http://localhost:8000
	testHttp()
}