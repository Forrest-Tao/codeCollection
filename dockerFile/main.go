package main

import (
	"flag"
	"fmt"
)

var fileName string

func init() {
	flag.StringVar(&fileName, "fileName", "defaultName", "enter FileName")
	flag.Parse()
	fmt.Println("in init function fileName:", fileName)
}

func main() {
	fmt.Println("in main function")

	fmt.Println(fileName)
	fmt.Println("done")
}
