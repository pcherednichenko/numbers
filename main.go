package main

import (
	"log"

	"github.com/pcherednichenko/numbers/app"
)

// main start server that combines numbers from url
func main() {
	bind := ":8080" // in feature better set bind in env of config
	server := app.New()
	log.Fatal(server.Start(bind))
}
