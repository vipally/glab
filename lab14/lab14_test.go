package lab14

import (
	"fmt"
	"math/big"
	"runtime"
	"sync"
	"testing"
	"time"
)

const (
	mps           = 20
	finishSecs    = 10
	readerCount   = 1000
	busyCount     = 50
	schduleInLock = true
)

var (
	frame int
	mutex sync.RWMutex
	wg    sync.WaitGroup
	chs   [readerCount]chan int
)

func TestMutex1WnR(t *testing.T) {
	start := time.Now()
	fmt.Println(start, "TestMutex1WnR start")
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
	for i := 1; i <= 100; i++ {
		if i == 50 {
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

func TestChannel1WnR(t *testing.T) {
	start := time.Now()
	fmt.Println(start, "TestChannel1WnR start")
	wg.Add(1 + readerCount)
	go chW()
	for i := 0; i < readerCount; i++ {
		go chR(i)
	}
	for i := 1; i <= busyCount; i++ {
		go busy(i)
	}
	wg.Wait()
	now := time.Now()
	fmt.Println(now, "TestChannel1WnR fnish,cost ", finishSecs, "->", now.Sub(start))
}

func chW() {
	for i := len(chs) - 1; i >= 0; i-- {
		chs[i] = make(chan int, 2)
	}
	t := time.NewTicker(time.Second / mps)
	fmsg := func(f int) {
		for i := len(chs) - 1; i >= 0; i-- {
			chs[i] <- f
		}
	}
	for {
		select {
		case <-t.C:
			frame++
			if schduleInLock && frame%mps == 0 {
				costLongTimeAndGosched()
				fmsg(frame)
			}
			if frame == finishSecs*mps {
				fmsg(frame)
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
