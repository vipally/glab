//针对1写N读的常见多线程应用案例
//分别应用mutext加锁和channel的方法 做多线程同步
//测试运行效率

package main

import (
	"fmt"
	"math/big"
	"runtime"
	"sync"
	"time"
)

const (
	mps           = 20
	finishSecs    = 600
	readerCount   = 5000
	busyCount     = 50
	schduleInLock = true
)

func init() {
	for i := len(chs) - 1; i >= 0; i-- {
		chs[i] = make(chan int, 2)
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
}

//chLen=2
// TestMutex1WnR start time=2018-11-28 09:02:36, finishSecs=300 readerCount=2000,busyCount=50
// 2018-11-28 09:07:59.5806564 +0800 CST m=+323.000474601 TestMutex1WnR fnish,cost 300 -> 5m22.9914741s
// TestMutex1WnR start time=2018-11-28 09:07:59, finishSecs=300 readerCount=2000,busyCount=50
// 2018-11-28 09:13:00.9508938 +0800 CST m=+624.370712001 TestChannel1WnR fnish,cost  300 -> 5m1.2482304s

//chLen=1
// TestMutex1WnR start time=2018-11-27 21:09:34, finishSecs=300 readerCount=2000,busyCount=50
// 2018-11-27 21:15:05.8786637 +0800 CST m=+331.694971901 TestMutex1WnR fnish,cost  300 -> 5m31.6899716s
// TestMutex1WnR start time=2018-11-27 21:15:05, finishSecs=300 readerCount=2000,busyCount=50
// 2018-11-27 21:20:44.4450285 +0800 CST m=+670.261336701 TestChannel1WnR fnish,cost  300 -> 5m38.5663648s
func main() {
	start := time.Now()
	costLongTimeAndGosched()
	now := time.Now()
	fmt.Println(start, now, "costLongTimeAndGosched fnish,cost ", now.Sub(start))
	Mutex1WnR()
	Channel1WnR()
}

var (
	frame int
	mutex sync.RWMutex
	wg    sync.WaitGroup
	chs   [readerCount]chan int
)

func Mutex1WnR() {
	//return
	frame = 0
	start := time.Now()
	fmt.Printf("TestMutex1WnR start time=%s, finishSecs=%d readerCount=%d,busyCount=%d\n", start.Format("2006-01-02 15:04:05"), finishSecs, readerCount, busyCount)
	wg.Add(1 + readerCount)
	go w()
	for i := 1; i <= readerCount; i++ {
		go r(i)
	}
	for i := 1; i <= busyCount; i++ {
		go busy(i)
	}
	wg.Wait()
	now := time.Now()
	fmt.Println(now, "TestMutex1WnR fnish,cost ", finishSecs, "->", now.Sub(start))
}

func w() {
	t := time.NewTicker(time.Second / mps)
	for {
		select {
		case <-t.C:
			mutex.Lock()
			frame++
			if schduleInLock && frame%mps == 0 {
				costLongTimeAndGosched()
			}
			if frame == finishSecs*mps {
				mutex.Unlock()
				wg.Done()
				return
			}
			mutex.Unlock()
		}
	}
}

func r(id int) {
	t := time.NewTicker(time.Second / mps)
	for {
		select {
		case <-t.C:
			mutex.RLock()
			if schduleInLock && frame%mps == 0 {
				costLongTimeAndGosched()
			}
			if frame == finishSecs*mps {
				mutex.RUnlock()
				wg.Done()
				return
			}
			mutex.RUnlock()
		}
	}
	//fmt.Println(time.Now(), id, "Rend")
}

func costLongTimeAndGosched() {
	b := big.NewInt(100)
	const count = 2000
	for i := 1; i <= count; i++ {
		if i == count/2 {
			runtime.Gosched()
		}
		b.MulRange(int64(i), int64(i+1))
	}
}

func busy(id int) {
	b := big.NewInt(100)
	for i := 1; i <= 10000000; i++ {
		runtime.Gosched()
		for j := 1; j <= 500; j++ {
			b.MulRange(int64(i), int64(j))
		}
	}
	fmt.Println(time.Now(), "busy end", id)
}

///////////////////////////////////////////////////////////////
//channel

func Channel1WnR() {
	frame = 0
	start := time.Now()
	fmt.Printf("TestMutex1WnR start time=%s, finishSecs=%d readerCount=%d,busyCount=%d\n", start.Format("2006-01-02 15:04:05"), finishSecs, readerCount, busyCount)
	wg.Add(1 + readerCount)
	for i := 0; i < readerCount; i++ {
		go chR(i)
	}
	go chW()
	for i := 1; i <= busyCount; i++ {
		go busy(i)
	}
	wg.Wait()
	now := time.Now()
	fmt.Println(now, "TestChannel1WnR fnish,cost ", finishSecs, "->", now.Sub(start))
}

func chW() {
	t := time.NewTicker(time.Second / mps)
	fmsg := func(f int) {
		//fmt.Println("fmsg", f)
		for i := len(chs) - 1; i >= 0; i-- {
			chs[i] <- f
		}
	}
	for {
		select {
		case <-t.C:
			frame++
			if schduleInLock && frame%mps == 0 {
				fmsg(frame)
				costLongTimeAndGosched()
			}
			if frame == finishSecs*mps {
				wg.Done()
				return
			}
		}
	}
}

func chR(id int) {
	for {
		select {
		case f := <-chs[id]:
			//fmt.Println("chR", id, f)
			if schduleInLock && f%mps == 0 {
				costLongTimeAndGosched()
			}
			if f == finishSecs*mps {
				wg.Done()
				return
			}
		}
	}
}
