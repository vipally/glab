package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var x int32
var wg sync.WaitGroup

func TestAtomic(t *testing.T) {
	wg.Add(2)
	go tt(1)
	go tt(2)
	wg.Wait()
	/*
		结论： atomic不会发送协程切换
			2 begin
			2 get 0
			2 add 1
			1 begin
			1 get 1
			1 add 2
	*/

}

func tt(i int) {
	fmt.Println(i, "begin")

	xx := atomic.LoadInt32(&x)

	fmt.Println(i, "get", xx)

	xy := atomic.AddInt32(&x, 1)

	fmt.Println(i, "add", xy)
	wg.Done()
}
