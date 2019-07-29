package main

import (
	"runtime"
	"sync"
	"time"
)

const workerNum = 100000
const outerLoop = 1000
const innerLoop = 10000

var wg sync.WaitGroup

func work(id int) {
	check := workerNum
	sum := 0
	start := time.Now()
	last := time.Now()
	for i := 1; i <= outerLoop; i++ {
		for j := 1; j <= innerLoop; j++ {
			sum += i * j
			// if id != check {
			// 	if j%1000 == 0 {
			// 		//time.Sleep(time.Microsecond)
			// 		runtime.Gosched()
			// 	}
			// }
		}
		if id == check {
			now := time.Now()
			println(i, now.String(), now.Sub(last).String())
			last = now
		}
		runtime.Gosched()
	}
	if id == check {
		println("finish, cost =", time.Now().Sub(start).String())
	}
	wg.Done()

}

func main() {
	start := time.Now()
	for i := 1; i <= workerNum; i++ {
		wg.Add(1)
		go work(i)
	}
	wg.Wait()
	println("all finish, cost =", time.Now().Sub(start).String())
}
