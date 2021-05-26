package main

import (
	"fmt"
	"math"
	"strconv"
)

// 默认值
func init() {
	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	println(defaultInt, defaultBool, defaultFloat32, defaultFloat64, defaultString)
}

// 整数数字
func int_type() {
	// 公式：[-2^(n-1), 2^(n-1) - 1]
	// range: -128 ~ 127
	var integer8 int8 = 127
	// range: -32768 ~ 32767
	var integer16 int16 = 32767
	// range: -2147483648 ~ 2147483647
	var integer32 int32 = 2147483647
	// range: -9223372036854775808 ~ 9223372036854775807
	var integer64 int64 = 9223372036854775807
	println(integer8, integer16, integer32, integer64)
}

// 无符号整数数字
func uint_type() {
	// 公式：[0, 2^n - 1]
	// range: 0 ~ 255
	var uinteger8 uint8 = 255
	// range: 0 ~ 65535
	var uinteger16 uint16 = 65535
	// range: 0 ~ 4294967295
	var uinteger32 uint32 = 4294967295
	// range: 0 ~ 18446744073709551615
	var uinteger64 uint64 = 18446744073709551615
	
	fmt.Println(uinteger8, uinteger16, uinteger32, uinteger64)
}

// 浮点数字
func float_type() {
	var float32 float32 = 2147483647
	var float64 float64 = 9223372036854775807
	println(float32, float64)

	println(math.MaxFloat32, math.MaxFloat64)
}

// 字符串
func str_type() {
	// \n：新行
	// \r：回车符
	// \t：选项卡
	// \'：单引号
	// \"：双引号
	// \\：反斜杠
	var firstName string = "John"
	lastName := "Doe"
	fullName := "John Doe \t(alias \"Foo\")\n"
	println(firstName, lastName, fullName)
}

// 布尔型
func bool_type() {
	var featureFlag bool = true
	println(featureFlag)
}

// 默认值
func default() {
	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	println(defaultInt, defaultBool, defaultFloat32, defaultFloat64, defaultString)
}

// 类型转换
func type_convert() {
	var integer16 int16 = 127
	var integer32 int32 = 32767
	println(int32(integer16), int32(integer16) + integer32)

	// string 转换为 int
	s2i, _ := strconv.Atoi("-42")
	// int 转换为 string
    i2s := strconv.Itoa(-42)
    println(s2i, i2s)
}

func main() {
	//int_type()
	//uint_type()
	//float_type()
	//bool_type()
	type_convert()

	// Unicode 字符（或 Unicode 码位）
	rune := 'G'
	println(rune)

	//var integer uint = -10
    //println(integer)
}
