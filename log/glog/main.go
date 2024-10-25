package main

import (
	"flag"
	"github.com/golang/glog"
)

// j
// go run main.go --stderrthreshold=INFO --log_dir=./
func main() {
	flag.Parse()
	defer glog.Flush()
	glog.Info("info")
	glog.Error("error")
	glog.Warning("warning")
	glog.Fatal("fatal")
}
