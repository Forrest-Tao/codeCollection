package main

import (
	"go.uber.org/ratelimit"
	"log"
	"os"
)

var file *os.File
var err error

func init() {
	file, err = os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.SetOutput(file)
}

func main() {
	r := ratelimit.New(10)
	log.Printf("----------in----------\n")
	r.Take()
	log.Println(" get one")

	for i := 0; i < 10; i++ {
		r.Take()
		log.Println(" get one")
	}
	log.Printf("----------out----------\n")
}
