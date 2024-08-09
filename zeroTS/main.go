package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()

	// 获取当前日期的零点零分时间
	zeroTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// 获取零点零分时间的 Unix 时间戳
	zeroTimestamp := zeroTime.Unix()

	fmt.Printf("当前时间的零点零分时间戳: %d\n", zeroTimestamp)
}
