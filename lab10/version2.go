package main

import (
	"fmt"
	"time"
)

type Version2 struct {
	maxDepth  int
	fnCallCnt int
}

func (v *Version2) R(n, depth int) (percent float64) {
	if depth == 1 {
		v.fnCallCnt = 0
		v.maxDepth = 1
	}
	if depth > v.maxDepth {
		v.maxDepth = depth
	}
	v.fnCallCnt++

	sum := float64(0)
	for x := int(1); x <= n; x++ {
		s := v.s(n, x, depth)
		sum += s
	}
	return sum / float64(n)
}

func (v *Version2) s(n, x, depth int) (percent float64) {
	switch {
	case x == 1:
		return 100
	case x == n:
		return 0
	default:
		return v.R(n-x+1, depth+1)
	}
}

func main2() {
	NN := 50
	for N := 1; N <= NN; N++ {
		v := &Version2{}
		start := time.Now()
		rate := v.R(N, 1)
		dur := time.Now().Sub(start) / time.Millisecond * time.Millisecond
		fmt.Printf("v2 N=%-2d R=%-5.1f T=%-10s RD=%-2d Fc=%d\n", N, rate, dur, v.maxDepth, v.fnCallCnt)
	}
	// output:
	// v2 N=1  R=100.0 T=0s         RD=1  RN=1
	// v2 N=2  R=50.0  T=0s         RD=1  RN=2
	// v2 N=3  R=50.0  T=0s         RD=2  RN=5
	// v2 N=4  R=50.0  T=0s         RD=3  RN=11
	// v2 N=5  R=50.0  T=0s         RD=4  RN=23
	// v2 N=6  R=50.0  T=0s         RD=5  RN=47
	// v2 N=7  R=50.0  T=0s         RD=6  RN=95
	// v2 N=8  R=50.0  T=0s         RD=7  RN=191
	// v2 N=9  R=50.0  T=0s         RD=8  RN=383
	// v2 N=10 R=50.0  T=0s         RD=9  RN=767
	// v2 N=11 R=50.0  T=0s         RD=10 RN=1535
	// v2 N=12 R=50.0  T=0s         RD=11 RN=3071
	// v2 N=13 R=50.0  T=0s         RD=12 RN=6143
	// v2 N=14 R=50.0  T=0s         RD=13 RN=12287
	// v2 N=15 R=50.0  T=0s         RD=14 RN=24575
	// v2 N=16 R=50.0  T=1ms        RD=15 RN=49151
	// v2 N=17 R=50.0  T=2ms        RD=16 RN=98303
	// v2 N=18 R=50.0  T=3ms        RD=17 RN=196607
	// v2 N=19 R=50.0  T=6ms        RD=18 RN=393215
	// v2 N=20 R=50.0  T=13ms       RD=19 RN=786431
	// v2 N=21 R=50.0  T=24ms       RD=20 RN=1572863
	// v2 N=22 R=50.0  T=30ms       RD=21 RN=3145727
	// v2 N=23 R=50.0  T=40ms       RD=22 RN=6291455
	// v2 N=24 R=50.0  T=84ms       RD=23 RN=12582911
	// v2 N=25 R=50.0  T=167ms      RD=24 RN=25165823
	// v2 N=26 R=50.0  T=330ms      RD=25 RN=50331647
	// v2 N=27 R=50.0  T=668ms      RD=26 RN=100663295
	// v2 N=28 R=50.0  T=1.371s     RD=27 RN=201326591
	// v2 N=29 R=50.0  T=2.677s     RD=28 RN=402653183
	// v2 N=30 R=50.0  T=5.292s     RD=29 RN=805306367
	// v2 N=31 R=50.0  T=10.769s    RD=30 RN=1610612735
	// v2 N=32 R=50.0  T=21.882s    RD=31 RN=3221225471
	// v2 N=33 R=50.0  T=45.926s    RD=32 RN=6442450943
	// v2 N=34 R=50.0  T=1m29.567s  RD=33 RN=12884901887
	// v2 N=35 R=50.0  T=2m55.502s  RD=34 RN=25769803775
	// v2 N=36 R=50.0  T=5m42.473s  RD=35 RN=51539607551
	// ...
}
