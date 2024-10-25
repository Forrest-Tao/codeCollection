package main

import (
	"log"
	"os"
)

var err error
var file *os.File

func init() {
	file, err = os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	log.SetOutput(file)
}

func cleanup() {
	defer file.Close()
}
