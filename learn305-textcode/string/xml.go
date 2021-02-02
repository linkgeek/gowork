package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	//给结构体赋值
	p := person{Name: "霜花似雪", Age: 18}

	var data []byte
	var err error

	//将结构体类型序列化出来，以xml格式展现,
	//MarshalIndent(p,""," "), 第二个参数是前缀，第三个参数是缩进
	if data, err = xml.MarshalIndent(p, "", " "); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	p2 := new(person)

	//反序列化
	if err = xml.Unmarshal(data, p2); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p2)
}
