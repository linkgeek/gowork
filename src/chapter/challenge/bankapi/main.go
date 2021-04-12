package main

import (
    "fmt"
    "net/http"
    "github.com/msft/bank"
    "strconv"
    "log"
    "encoding/json"
)

var accounts = map[float64]*CustomAccount{}

// 公开对账单方法 http://localhost:8000/statement?number=1001
func statement0(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")

    if numberqs == "" {
        fmt.Fprintf(w, "Account number is missing!")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account number!")
    } else {
        account, ok := accounts[number]
        if !ok {
            fmt.Fprintf(w, "Account with number %v can't be found!", number)
        } else {
            fmt.Fprintf(w, account.Statement())
        }
    }
}

func statement(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")

    if numberqs == "" {
        fmt.Fprintf(w, "Account number is missing!")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account number!")
    } else {
        if account, ok := accounts[number]; !ok {
            fmt.Fprintf(w, "Account with number %v can't be found!", number)
        } else {
            json.NewEncoder(w).Encode(bank.Statement(account))
        }
    }
}

// 公开存款方法 
// http://localhost:8000/deposit?number=1001&amount=10
// curl http://localhost:8000/deposit?number=1001&amount=100
func deposit(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")
    amountqs := req.URL.Query().Get("amount")

    if numberqs == "" {
        fmt.Fprintf(w, "Account number is missing!")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account number!")
    } else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid amount number!")
    } else {
        account, ok := accounts[number]
        if !ok {
            fmt.Fprintf(w, "Account with number %v can't be found!", number)
        } else {
            err := account.Deposit(amount)
            if err != nil {
                fmt.Fprintf(w, "%v", err)
            } else {
                fmt.Fprintf(w, account.Statement())
            }
        }
    }
}

// 公开取款方法 http://localhost:8000/withdraw?number=1001&amount=100
func withdraw(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")
    amountqs := req.URL.Query().Get("amount")

    if numberqs == "" {
        fmt.Fprintf(w, "Account number is missing!")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account number!")
    } else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid amount number!")
    } else {
        account, ok := accounts[number]
        if !ok {
            fmt.Fprintf(w, "Account with number %v can't be found!", number)
        } else {
            err := account.WithDraw(amount)
            if err != nil {
                fmt.Fprintf(w, "%v", err)
            } else {
                fmt.Fprintf(w, account.Statement())
            }
        }
    }
}

// 实现转账方法 http://localhost:8000/transfer?number=1001&dest=1002&amount=100
func transfer(w http.ResponseWriter, req *http.Request) {
    numberqs := req.URL.Query().Get("number")
    destqs := req.URL.Query().Get("dest")
    amountqs := req.URL.Query().Get("amount")

    if numberqs == "" {
        fmt.Fprintf(w, "Account number is missing!")
        return
    }

    if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account number!")
    } else if dest, err := strconv.ParseFloat(destqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid account destination number!")
    } else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
        fmt.Fprintf(w, "Invalid amount number!")
    } else {
        if accountA, ok := accounts[number]; !ok {
            fmt.Fprintf(w, "Account with number %v can't be found!", number)
        } else if accountB, ok := accounts[dest]; !ok {
            fmt.Fprintf(w, "Account with number %v can't be found!", dest)
        } else {
            err := accountA.Transfer(amount, accountB.Account)
            if err != nil {
                fmt.Fprintf(w, "%v", err)
            } else {
                fmt.Fprintf(w, accountA.Statement())
            }
        }
    }
}

type CustomAccount struct {
    *bank.Account
}

func (c *CustomAccount) Statement() string {
    json, err := json.Marshal(c)
    if err != nil {
        return err.Error()
    }
    return string(json)
}

func main() {
    accounts[1001] = &CustomAccount{
        Account: &bank.Account{
            Customer: bank.Customer{
                Name: "John1",
                Address: "ShenZhen",
                Phone: "183xxxx",
            },
            Number: 1001,
        },
    }
    accounts[1002] = &CustomAccount{
        Account: &bank.Account{
            Customer: bank.Customer{
                Name: "John2",
                Address: "ShenZhen",
                Phone: "199xxxx",
            },
            Number: 1002,
        },
    }

    http.HandleFunc("/statement", statement)
    http.HandleFunc("/deposit", deposit)
    http.HandleFunc("/withdraw", withdraw)
    http.HandleFunc("/transfer", transfer)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}