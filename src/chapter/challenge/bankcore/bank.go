package bank

import (
    "errors"
    "fmt"
)

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
func (a *Account) Statement0() string {
    return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

// 实现转账方法
func (a *Account) Transfer(amount float64, dest *Account) error {
    if amount <= 0 {
        return errors.New("the amount to transfer  should be greater than zero")
    }

    if a.Balance < amount {
        return errors.New("the amount to transfer should be greater than the account's balance")
    }

    a.WithDraw(amount)
    dest.Deposit(amount)
    return nil
}

type Bank interface {
    Statement() string
}

func Statement(b Bank) string {
    return b.Statement()
}