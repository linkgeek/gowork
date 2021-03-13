package model

import "fmt"

type AliPay struct {
	PaymentArgs
	AliOpenID string
}

func (this *AliPay) Info() {
	fmt.Printf("AliPay = %v\n", this)
}
