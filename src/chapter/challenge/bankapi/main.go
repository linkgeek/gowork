package main

import (
    "fmt"
    "net/http"
    "github.com/msft/bank"
    "strconv"
    "log"
)

var accounts = map[float64]*bank.Account{}

// 公开对账单方法 http://localhost:8000/statement?number=1001
func statement(w http.ResponseWriter, req *http.Request) {
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

func main() {
    accounts[1001] = &bank.Account{
        Customer: bank.Customer{
            Name: "John",
            Address: "ShenZhen",
            Phone: "183xxxx",
        },
        Number: 1001,
    }

    http.HandleFunc("/statement", statement)
    http.HandleFunc("/deposit", deposit)
    http.HandleFunc("/withdraw", withdraw)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}