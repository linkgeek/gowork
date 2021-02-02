package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

//获取xml格式的节点属性值
func getAttributeValue(attr []xml.Attr, name string) string {
	for _, a := range attr {
		if a.Name.Local == name {
			return a.Value
		}
	}

	return ""
}

func main() {
	//读取文件
	content, err := ioutil.ReadFile("vsproj.csproj")
	//解析
	docoder := xml.NewDecoder(bytes.NewBuffer(content))

	var t xml.Token
	var inItemGroup bool
	for t, err = docoder.Token(); err == nil; t, err = docoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			//fmt.Println(name) //打印出xml的开始节点
			if inItemGroup {
				if name == "Compile" {
					fmt.Println(getAttributeValue(token.Attr, "Include"))
				}
			} else {
				if name == "ItemGroup" {
					inItemGroup = true
				}
			}
		case xml.EndElement:
			if inItemGroup {
				if token.Name.Local == "ItemGroup" {
					inItemGroup = false
				}
			}
		}
	}
}
