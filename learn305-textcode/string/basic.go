package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串基本操作
func str_basic() {
	s := "hello world!"

	// 是否包含
	fmt.Println(strings.Contains(s, "hello"), strings.Contains(s, "?"))

	// 索引，base 0
	fmt.Println(strings.Index(s, "o"))

	ss := "1#2#345"

	// 切割字符串
	splitedStr := strings.Split(ss, "#")
	fmt.Println(splitedStr)

	// 合并字符串
	fmt.Println(strings.Join(splitedStr, "#"))

	// 前缀、后缀
	fmt.Println(strings.HasPrefix(s, "he"), strings.HasSuffix(s, "ld"))
}

// 字符串转换
func str_convert() {
	// 整型与字符串转换
	fmt.Println(strconv.Itoa(10))
	fmt.Println(strconv.Atoi("711"))

	// 解析
	fmt.Println(strconv.ParseBool("false"))
	fmt.Println(strconv.ParseFloat("3.14", 64))

	// 格式化
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(123, 2))
}

func main() {
	//str_basic()
	str_convert()
}
