package bank

import (
    "errors"
    "fmt"
)

func Hello() string {
    return "Hey Cat! I'm working!"
}

type Customer struct {
    Name string
    Address string
    Phone string
}

type Account struct {
    Customer
    Number int32
    Balance float64
}

// 实现存款方法
func (a *Account) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("the amount to deposit should be greater than zero")
    }

    a.Balance += amount
    return nil
}

// 实现取款方法
func (a *Account) WithDraw(amount float64) error {
    if amount <= 0 {
        return errors.New("the amount to deposit should be greater than zero")
    }

    if a.Balance < amount {
        return errors.New("the amount to withdraw should be greater than the account's balance")
    }

    a.Balance -= amount
    return nil
}

// 实现对账单方法
func (a *Account) Statement() string {
    return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}