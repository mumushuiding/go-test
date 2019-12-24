package test

import (
	"fmt"
	"testing"
	"time"
)

// 每5秒执行一次
// func TestCycle(t *testing.T) {
// 	fmt.Println("===================")
// 	ticker := time.NewTicker(time.Second * 5)
// 	go func() {
// 		for _ = range ticker.C {
// 			log.Println(time.Now())
// 		}
// 	}()
// 	time.Sleep(time.Minute)
// }

// TestCron1 每天0时执行
func TestCron1(t *testing.T) {
	go func() {
		for {
			now := time.Now()
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			fmt.Println(next)
			t := time.NewTimer(next.Sub(now))
			<-t.C
			fmt.Println(time.Now())
		}
	}()
	time.Sleep(time.Minute)
}
