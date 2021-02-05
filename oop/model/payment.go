package model

import "fmt"

type PaymentArgs struct {
	AppID string
	MchID string
	Token string
}

func (this *PaymentArgs) Info() {
	fmt.Printf("Info = %v\n", this)
}
