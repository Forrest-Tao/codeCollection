package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("in demp")
	writer := bufio.NewWriter(os.Stdout)
	io.Copy(writer, reader)
	writer.Flush()
}
