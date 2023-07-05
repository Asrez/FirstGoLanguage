package main

import (
	// "github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// logging
		log.Print("Test")
		log.Info().Str("Category" , "search").Msg("Searching...")
}

// trace
// debug
// info 
// warn 
// error
// fatal
// panic
