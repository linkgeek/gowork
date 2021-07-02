package model

type userInfo struct {
	Name   string
	Age    int
	Height float32
}

func NewUserInfo(name string, age int, height float32) *userInfo {
	return &userInfo{
		Name:   name,
		Age:    age,
		Height: height,
	}
}
