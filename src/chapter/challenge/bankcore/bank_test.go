package bank

import "testing"

// 测试创建账号
func TestAccount(t *testing.T) {
	account := Account{
        Customer: Customer{
            Name:    "John",
            Address: "Los Angeles, California",
            Phone:   "(213) 555 0147",
        },
        Number:  1001,
        Balance: 0,
    }

    if account.Name == "" {
        t.Error("can't create an Account object")
    }
}

// 测试存款方法
func TestDeposit(t *testing.T) {
	account := Account{
        Customer: Customer{
            Name:    "John",
            Address: "Los Angeles, California",
            Phone:   "(213) 555 0147",
        },
        Number:  1001,
        Balance: 0,
	}
	
	account.Deposit(10)

    if account.Balance != 10 {
        t.Error("balance is not being updated after a deposit")
    }
}

// 测试负数存款额
func TestDepositInvalid(t *testing.T) {
	account := Account{
        Customer: Customer{
            Name:    "John",
            Address: "Los Angeles, California",
            Phone:   "(213) 555 0147",
        },
        Number:  1001,
        Balance: 0,
	}

    if err := account.Deposit(-10); err == nil {
        t.Error("only positive numbers should be allowed to deposit")
    }
}

// 测试取款方法
func TestWithDraw(t *testing.T)  {
    account := Account{
        Customer: Customer{
            Name:    "John",
            Address: "Los Angeles, California",
            Phone:   "(213) 555 0147",
        },
        Number:  1001,
        Balance: 0,
	}
	
	account.Deposit(10)
	account.WithDraw(10)

    if account.Balance != 0 {
        t.Error("balance is not being updated after withdraw")
    }
}

// 测试对账单方法
func TestStatement(t *testing.T)  {
    account := Account{
        Customer: Customer{
            Name:    "John",
            Address: "Los Angeles, California",
            Phone:   "(213) 555 0147",
        },
        Number:  1001,
        Balance: 0,
	}
	
	account.Deposit(100)
	statement := account.Statement()
    if statement != "1001 - John - 100" {
        t.Error("statement doesn't have the proper format")
    }
}

// 测试转账方法
func TestTransfer(t *testing.T)  {
    accountA := Account{
        Customer: Customer{
            Name:    "John",
            Address: "Los Angeles, California",
            Phone:   "(213) 555 0147",
        },
        Number:  1001,
        Balance: 0,
	}

    accountB := Account{
        Customer: Customer{
            Name:    "Mark",
            Address: "Los Angeles, California",
            Phone:   "(213) 555 0147",
        },
        Number:  1002,
        Balance: 0,
	}
	
	accountA.Deposit(100)
    err := accountA.Transfer(50, &accountB)

    if accountA.Balance != 50 && accountB.Balance != 50 {
        t.Error("transfer from account A to account B is not working", err)
    }
}

// go test -v