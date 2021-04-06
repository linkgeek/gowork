package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func testZero()  {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("hello world")

	log.Debug().
        Int("EmployeeID", 1001).
        Msg("Getting employee information")

    log.Debug().
        Str("Name", "John").
        Send()
}

func main() {
	testZero()
}