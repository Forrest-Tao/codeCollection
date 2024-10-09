package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	nodeName := os.Getenv("NODE_NUMBER")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "This pod is scheduled on node %s\n", nodeName)
	})

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
