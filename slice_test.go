package test

import (
	"fmt"
	"testing"
)

var concurrency []int
var channel chan int

// func TestSliceRemove(t *testing.T) { // slice删除指定位置的元素
// 	a := []int{1, 2, 3, 4, 5}
// 	// 删除第二个
// 	a = append(a[:1], a[2:]...)
// 	fmt.Println(a)
// }
func TestSliceAddConcurrently(t *testing.T) {
	channel = make(chan int, 1000)
	for i := 0; i < 1000; i++ {
		go addToChanConcurrently(i)
	}
	for j := 0; j < 1000; j++ {
		add(<-channel)
	}
	close(channel)
	for _, x := range concurrency {
		fmt.Println(x)
	}
	fmt.Printf("共有元素个数: %d", len(concurrency))
}
func addToChanConcurrently(a int) {
	channel <- a
}
func add(a int) {
	concurrency = append(concurrency, a)
}
