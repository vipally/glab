package lab15

import (
	"fmt"
	"testing"
)

const (
	randCount = 1000000
)

// little   p=0 count=486801/1000000 rate=48.68%
// big      p=0 count=485346/1000000 rate=48.53%
// baozi    p=1 count=4835/1000000 rate=0.48%
// baozi    p=2 count=4602/1000000 rate=0.46%
// baozi    p=3 count=4444/1000000 rate=0.44%
// baozi    p=4 count=4430/1000000 rate=0.44%
// baozi    p=5 count=4662/1000000 rate=0.46%
// baozi    p=6 count=4880/1000000 rate=0.48%
// point    p=3 count=4835/1000000 rate=0.48%
// point    p=4 count=14069/1000000 rate=1.40%
// point    p=5 count=28123/1000000 rate=2.81%
// point    p=6 count=46551/1000000 rate=4.65%
// point    p=7 count=69324/1000000 rate=6.93%
// point    p=8 count=97853/1000000 rate=9.78%
// point    p=9 count=115485/1000000 rate=11.54%
// point    p=10 count=124442/1000000 rate=12.44%
// point    p=11 count=124379/1000000 rate=12.43%
// point    p=12 count=115953/1000000 rate=11.59%
// point    p=13 count=96746/1000000 rate=9.67%
// point    p=14 count=69012/1000000 rate=6.90%
// point    p=15 count=46388/1000000 rate=4.63%
// point    p=16 count=27798/1000000 rate=2.77%
// point    p=17 count=14162/1000000 rate=1.41%
// point    p=18 count=4880/1000000 rate=0.48%
func TestRand(t *testing.T) {
	var (
		little = 0
		big    = 0
		baozi  [6]int
		point  [16]int
	)
	for i := 0; i < randCount; i++ {
		_, p, bz := randDice()
		if !bz {
			if p <= 10 {
				little++
			} else {
				big++
			}
		} else {
			baozi[(p/3)-1]++
		}
		point[p-3]++
	}
	showResult("little", 0, little)
	showResult("big", 0, big)
	for i, v := range baozi {
		showResult("baozi", 1+i, v)
	}
	for i, v := range point {
		showResult("point", 3+i, v)
	}
}

func showResult(name string, point int, count int) {
	rate := count * 10000 / randCount
	fmt.Printf("%-8s p=%d count=%d/%d rate=%d.%d%%\n", name, point, count, randCount, rate/100, rate%100)
}
