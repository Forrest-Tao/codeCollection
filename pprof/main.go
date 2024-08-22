package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	r := gin.Default()
	go func() {
		http.ListenAndServe(":6060", nil) //pprof的位置
	}()
	r.GET("/readall", testReadAll)
	r.GET("/iocopy", testIOCopy)
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func testIOCopy(c *gin.Context) {
	//fmt.Println("in testIOCopy")
	file, err := os.Open("./cka.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		panic(err)
	}
}

func testReadAll(c *gin.Context) {
	//fmt.Println("in testReadAll")
	file, err := os.Open("./cka.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//simulate onLine err
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	c.Writer.Write(bytes)
	if err != nil {
		panic(err)
	}
}
