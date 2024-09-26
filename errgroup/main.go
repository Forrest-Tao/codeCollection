package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	eg := errgroup.Group{}

	eg.Go(func() error {
		fmt.Println("begin func-1")
		fmt.Println("end func-1")
		return errors.New("err ...")
	})

	eg.Go(func() error {
		fmt.Println("begin func-2")
		time.Sleep(time.Second * 5)
		fmt.Println("end func-2")
		return nil
	})
	if err := eg.Wait(); err != nil {
		fmt.Println("err is not nil")
	}
	fmt.Println("done")
}
