package debuglock

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const maxFileLen = 20
const maxFuncLen = 20
const tooLongLock = time.Second

type DebugRWMutex struct {
	l   sync.RWMutex
	cid uint64 // callerId counter
	e   uint64 // epoch counter
	r   int32  // reader count
	w   int32  // writer count
	wr  int32  // wait reader count
	ww  int32  // wait writer count
}

func getGID() uint64 {
	if true { //do not use slow debug goroutine id
		b := make([]byte, 64)
		b = b[:runtime.Stack(b, false)]
		b = bytes.TrimPrefix(b, []byte("goroutine "))
		b = b[:bytes.IndexByte(b, ' ')]
		n, _ := strconv.ParseUint(string(b), 10, 64)
		return n
	}

	return 0
}

func (l *DebugRWMutex) log(fn string) func() {
	cid := atomic.AddUint64(&l.cid, 1)
	pc, file, line, _ := runtime.Caller(2)
	if len(file) > maxFileLen {
		file = file[len(file)-maxFileLen:]
	}

	cf := ""
	if p := runtime.FuncForPC(pc); p != nil {
		cf = p.Name()
		if len(cf) > maxFuncLen {
			cf = cf[len(cf)-maxFuncLen:]
		}
	}

	start := time.Now()
	gid := getGID()
	fmt.Printf("%s [dbg-lock-beg  %-2s] [%s:%-4d-%s] %- 8s cid=%-6d gid=%-6d p=%p e=%-6d r=%-4d wr=%-4d w=%-4d ww=%-4d\n",
		time.Now().Format("2006-01-02T15:04:05.000000"), fn, file, line, cf, "-", cid, gid, l, atomic.LoadUint64(&l.e),
		atomic.LoadInt32(&l.r), atomic.LoadInt32(&l.wr), atomic.LoadInt32(&l.w), atomic.LoadInt32(&l.ww))
	deferfun := func() {
		cost := time.Now().Sub(start) / time.Millisecond * time.Millisecond
		if cost > tooLongLock {
			fmt.Printf("%s [dbg-lock-enl %-2s] [%s:%-4d-%s] t=%- 6s cid=%-6d gid=%-6d p=%p e=%-6d r=%-4d wr=%-4d w=%-4d ww=%-4d\n",
				time.Now().Format("2006-01-02T15:04:05.000000"), fn, file, line, cf, cost, cid, gid, l, atomic.LoadUint64(&l.e),
				atomic.LoadInt32(&l.r), atomic.LoadInt32(&l.wr), atomic.LoadInt32(&l.w), atomic.LoadInt32(&l.ww))
		} else {
			fmt.Printf("%s [dbg-lock-end  %-2s] [%s:%-4d-%s] t=%- 6s cid=%-6d gid=%-6d p=%p e=%-6d r=%-4d wr=%-4d w=%-4d ww=%-4d\n",
				time.Now().Format("2006-01-02T15:04:05.000000"), fn, file, line, cf, cost, cid, gid, l, atomic.LoadUint64(&l.e),
				atomic.LoadInt32(&l.r), atomic.LoadInt32(&l.wr), atomic.LoadInt32(&l.w), atomic.LoadInt32(&l.ww))
		}
	}
	return deferfun
}

func (l *DebugRWMutex) Lock() {
	deferfun := l.log("L")
	defer deferfun()

	atomic.AddInt32(&l.ww, 1)
	l.l.Lock()
	atomic.AddInt32(&l.ww, -1)

	atomic.AddUint64(&l.e, 1)
	atomic.AddInt32(&l.w, 1)
}

func (l *DebugRWMutex) Unlock() {
	deferfun := l.log("U")
	defer deferfun()

	atomic.AddInt32(&l.w, -1)
	l.l.Unlock()
}

func (l *DebugRWMutex) RLock() {
	deferfun := l.log("RL")
	defer deferfun()

	atomic.AddInt32(&l.wr, 1)
	l.l.RLock()
	atomic.AddInt32(&l.wr, -1)

	if n := atomic.AddInt32(&l.r, 1); n == 1 {
		atomic.AddUint64(&l.e, 1) //first rlock ok,new epoch
	}
}

func (l *DebugRWMutex) RUnlock() {
	deferfun := l.log("RU")
	defer deferfun()

	atomic.AddInt32(&l.r, -1)
	l.l.RUnlock()
}
