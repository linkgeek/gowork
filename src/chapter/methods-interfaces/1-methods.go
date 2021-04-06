package main

import (
	"fmt"
	"strings"
)

type triangle struct {
    size int
}

// 嵌入方法
type coloredTriangle struct {
    triangle
    color string
}

type square struct {
    size int
}

func (t triangle) perimeter() int {
    return t.size * 3
}

// 重载方法
func (ct coloredTriangle) perimeter() int {
    return ct.size * 3 * 2
}

func (s square) perimeter() int {
    return s.size * 4
}

func (t *triangle) doubleSize() {
    t.size *= 2
}

type upperstring string

func (s upperstring) Upper() string {
    return strings.ToUpper(string(s))
}

func main() {
    t := triangle{3}
	t.doubleSize()
    s := square{4}
    fmt.Println("Perimeter (triangle):", t.perimeter())
    fmt.Println("Perimeter (square):", s.perimeter())

	sup := upperstring("Learning Go!")
    fmt.Println(sup)
    fmt.Println(sup.Upper())

	ct := coloredTriangle{triangle{3}, "red"}
	fmt.Println("Size:", ct.size)
    fmt.Println("Perimeter (colored):", ct.perimeter())
    fmt.Println("Perimeter (normal):", ct.triangle.perimeter())
}