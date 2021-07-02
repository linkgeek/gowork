package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    start := time.Now()

    apis := []string{
        "https://management.azure.com",
        "https://dev.azure.com",
        "https://api.github.com",
        "https://outlook.office.com/",
        "https://api.somewhereintheinternet.com/",
        "https://graph.microsoft.com",
    }

    for _, api := range apis {
        _, err := http.Get(api)
        if err != nil {
            fmt.Printf("ERROR: %s is down!\n", api)
            continue
        }

        fmt.Printf("SUCCESS: %s is up and running!\n", api)
    }

    elapsed := time.Since(start)
    fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}