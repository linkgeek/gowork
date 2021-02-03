package main

import (
	"fmt"
	"oop/model"
)

func main() {
	// user := model.userInfo{
	// 	Name:   "giant",
	// 	Age:    19,
	// 	Height: 173,
	// }

	user := model.NewUserInfo("giant", 17, 178)
	fmt.Print(user.Name, user.Age)
}
