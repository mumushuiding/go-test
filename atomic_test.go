package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var number uint32

func TestAtomic(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		// go addNumber(&wg)           // 结果不为 1000 并发不安全
		go addNumberWithAtomic(&wg) // 结果为1000 并发安全
	}
	wg.Wait()
	fmt.Printf("complete:%d", number)
}

// addNumberWithAtomic 并发安全
func addNumberWithAtomic(wg *sync.WaitGroup) {
	time.Sleep(time.Duration(1) * time.Second)
	atomic.AddUint32(&number, 1)
	fmt.Println(number)
	wg.Done()
}
func addNumber(wg *sync.WaitGroup) {
	time.Sleep(time.Duration(1) * time.Second)
	number++
	fmt.Println(number)
	wg.Done()
}
