//针对1写N读的常见多线程应用案例
//分别应用mutext加锁和channel的方法 做多线程同步
//测试运行效率
//从测试结果可以看出 加锁的方式执行效率较低
//这里设置了一个典型的测试点 每一秒会触发一次较耗时的操作costLongTimeAndGosched
//加锁方式下，这个操作相当于是串行执行，因此效率较低。
//而channel模式，共享的frame是通过channel通知到每个Reader线程，因此不存在资源竞争的情况，天然可并行
//这种模型也更符合go语言推荐的多线程模型。

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
	finishSecs    = 60
	readerCount   = 5000
	busyCount     = 50
	chLen         = 1
	schduleInLock = true
)

var (
	frame    int
	mutex    sync.RWMutex
	wg       sync.WaitGroup
	chs      [readerCount]chan int
	chWriter chan int
)

func init() {
	chWriter = make(chan int, 10)
	for i := len(chs) - 1; i >= 0; i-- {
		chs[i] = make(chan int, chLen)
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
	fmt.Println(start, now, "costLongTimeAndGosched finish,cost ", now.Sub(start))
	//Mutex1WnR()
	//Channel1WnR()
	ChannelnW1R()
}

func Mutex1WnR() {
	//return
	frame = 0
	start := time.Now()
	fmt.Printf("Mutex1WnR start time=%s, finishSecs=%d readerCount=%d,busyCount=%d\n", start.Format("2006-01-02 15:04:05"), finishSecs, readerCount, busyCount)
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
	fmt.Println(now, "Mutex1WnR finish,cost ", finishSecs, "->", now.Sub(start))
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
		count := 1000
		if i%100 == 0 {
			count = 5000
		}
		for j := 1; j <= count; j++ {
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
	fmt.Printf("Mutex1WnR start time=%s, finishSecs=%d readerCount=%d,busyCount=%d chLen=%d\n", start.Format("2006-01-02 15:04:05"), finishSecs, readerCount, busyCount, chLen)
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
	fmt.Println(now, "Channel1WnR finish,cost ", finishSecs, "->", now.Sub(start))
}

func chTellReaders(f int) {
	log := (f/mps)%10 == 0
	var start time.Time
	if log {
		start = time.Now()
		fmt.Println(start.Format("2006-01-02 15:04:05.999999999"), "chTellReaders", f)
	}
	for i := len(chs) - 1; i >= 0; i-- {
		//fmt.Println("chW", i)
		chs[i] <- f
		//fmt.Println("chW", i, f)
	}
	if log {
		now := time.Now()
		fmt.Println(now.Format("2006-01-02 15:04:05.999999999"), now.Sub(start), "chTellReaders end", f)
	}
}

func chW() {
	t := time.NewTicker(time.Second / mps)
	for {
		select {
		case <-t.C:
			frame++
			if schduleInLock && frame%mps == 0 {
				chTellReaders(frame)
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
		//fmt.Println("chR", id)
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

func ChannelnW1R() {
	start := time.Now()
	fmt.Printf("ChannelnW1R start time=%s, finishSecs=%d readerCount=%d,busyCount=%d chLen=%d\n", start.Format("2006-01-02 15:04:05"), finishSecs, readerCount, busyCount, chLen)
	wg.Add(1 + readerCount)
	go chMain()
	for i := 0; i < readerCount; i++ {
		go chWx(i)
	}
	for i := 1; i <= busyCount; i++ {
		go busy(i)
	}
	wg.Wait()
	now := time.Now()
	fmt.Println(now, "ChannelnW1R finish,cost ", finishSecs, "->", now.Sub(start))
}

func chWx(id int) {
	t := time.NewTicker(time.Second / mps)
	for count := finishSecs * mps; count > 0; count-- {
		select {
		case <-t.C:
			chWriter <- id
		}
	}
	wg.Done()
}

func chMain() {
	finishCount := 0
	var cnt [readerCount]int
	for {
		select {
		case id := <-chWriter:
			cnt[id]++
			c := cnt[id]
			if c == finishSecs*mps {
				if finishCount++; finishCount == readerCount {
					wg.Done()
					return
				}
			}
		}
	}
}
