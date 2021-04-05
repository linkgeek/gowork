package main

import (
    "log"
	"os"
	"github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func testOs()  {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    log.SetOutput(file)
    log.Print("Hey, I'm a log!")
}

func zlog()  {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
    log.Print("Hey! I'm a log message!")
}

func main() {
    zlog()
}