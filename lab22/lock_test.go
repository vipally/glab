package debuglock

import (
	"sync"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	var l DebugRWMutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		go func() {
			wg.Add(1)
			l.Lock()
			time.Sleep(time.Microsecond)
			l.Unlock()
			wg.Done()
		}()
	}
	for i := 0; i < 1000; i++ {
		go func() {
			wg.Add(1)
			l.RLock()
			time.Sleep(time.Microsecond)
			l.RUnlock()
			wg.Done()
		}()
	}
	wg.Wait()
}
