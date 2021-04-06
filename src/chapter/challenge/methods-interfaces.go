package challenge

import (
	"fmt"
	"errors"
)

type Account struct {
	FirstName string
	LastName string
}
// 更改名称
func (a *Account) ChangeName(newname string) {
	a.FirstName = newname
}

type Employee struct {
	Account
	Credits float64
}

// 打印名称
func (e Employee) String() string {
	return fmt.Sprintf("Name: %s %s\nCredits: %.2f\n", e.FirstName, e.LastName, e.Credits)
}

// 添加贷方
func (e *Employee) AddCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		e.Credits += amount
		return e.Credits, nil
	}
	return 0.0, errors.New("Invalid credit amount.")
}

// 删除贷方
func (e *Employee) RemoveCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		if amount <= e.Credits{
			e.Credits -= amount
			return e.Credits, nil
		}
		return 0.0, errors.New("You can't remove more credits than the account has.")
	}
	return 0.0, errors.New("You can't remove negative numbers.")
}

// 检查余额
func (e *Employee) checkCredits() float64 {
	return e.Credits
}