package main

import (
	//"runtime"
	"sync"
	"time"
)

const workerNum = 1000000000
const outerLoop = 1000
const innerLoop = 10000

var wg sync.WaitGroup

func work(id int) {
	sum := 0
	start := time.Now()
	last := time.Now()
	for i := 1; i <= outerLoop; i++ {
		for j := 1; j <= innerLoop; j++ {
			sum += i * j
			if id != 1 {
				if j%100 == 0 {
					time.Sleep(time.Microsecond)
				}
			}
		}
		if id == 1 {
			now := time.Now()
			println(i, now.String(), now.Sub(last).String())
			last = now
		}
		//runtime.Gosched()
	}
	if id == 1 {
		println("finish, cost =", time.Now().Sub(start).String())
	}
	wg.Done()
	println("all finish, cost =", time.Now().Sub(start).String())
}

func main() {
	for i := 1; i <= workerNum; i++ {
		wg.Add(1)
		go work(i)
	}
	wg.Wait()
}
