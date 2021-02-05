package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerID    int
	ServerPort  int
	ServerToken string
}

type Server2 struct {
	ServerID    int    `json:"id"`
	ServerPort  int    `json:"port"`
	ServerToken string `json:"token"`
}

// struct转json
func EncodeStruct() {
	serv := &Server{
		ServerID:    1,
		ServerPort:  80,
		ServerToken: "sdfsdf",
	}
	b, err := json.Marshal(serv)
	if err != nil {
		fmt.Println("Marshal err: ", err.Error())
		return
	}
	fmt.Println("Marshal json: ", string(b))
}

// map转json
func EncodeMap() {
	serv := make(map[string]interface{})
	serv["ServerID"] = 1
	serv["ServerPort"] = 8080
	serv["ServerToken"] = "srer"
	b, err := json.Marshal(serv)
	if err != nil {
		fmt.Println("Marshal err: ", err.Error())
		return
	}
	fmt.Println("Marshal json: ", string(b))
}

// 反序列化struct
func DecodeStruct() {
	jsonStr := `{"ServerID":1,"ServerPort":8080,"ServerToken":"srer"}`
	serv := new(Server)
	err := json.Unmarshal([]byte(jsonStr), &serv)
	if err != nil {
		fmt.Println("Unmarshal err: ", err.Error())
		return
	}
	fmt.Println("Unmarshal struct: ", serv)
}

// 反序列化map
func DecodeMap() {
	jsonStr := `{"ServerID":1,"ServerPort":8080,"ServerToken":"srer"}`
	serv := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &serv)
	if err != nil {
		fmt.Println("Unmarshal err: ", err.Error())
		return
	}
	fmt.Println("Unmarshal map: ", serv)
}

func main() {
	//EncodeStruct()
	//EncodeMap()

	DecodeStruct()
	DecodeMap()
}
