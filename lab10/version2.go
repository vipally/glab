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
	M := 50
	for N := 1; N <= M; N++ {
		v := &Version2{}
		start := time.Now()
		rate := v.R(N, 1)
		dur := time.Now().Sub(start) / time.Millisecond * time.Millisecond
		fmt.Printf("v2 N=%-2d R=%-5.1f T=%-10s RD=%-2d Fc=%d\n", N, rate, dur, v.maxDepth, v.fnCallCnt)
	}
	// output:
	// v2 N=1  R=100.0 T=0s         RD=1  Fc=1
	// v2 N=2  R=50.0  T=0s         RD=1  Fc=1
	// v2 N=3  R=50.0  T=0s         RD=2  Fc=2
	// v2 N=4  R=50.0  T=0s         RD=3  Fc=4
	// v2 N=5  R=50.0  T=0s         RD=4  Fc=8
	// v2 N=6  R=50.0  T=0s         RD=5  Fc=16
	// v2 N=7  R=50.0  T=0s         RD=6  Fc=32
	// v2 N=8  R=50.0  T=0s         RD=7  Fc=64
	// v2 N=9  R=50.0  T=0s         RD=8  Fc=128
	// v2 N=10 R=50.0  T=0s         RD=9  Fc=256
	// v2 N=11 R=50.0  T=0s         RD=10 Fc=512
	// v2 N=12 R=50.0  T=0s         RD=11 Fc=1024
	// v2 N=13 R=50.0  T=0s         RD=12 Fc=2048
	// v2 N=14 R=50.0  T=0s         RD=13 Fc=4096
	// v2 N=15 R=50.0  T=0s         RD=14 Fc=8192
	// v2 N=16 R=50.0  T=0s         RD=15 Fc=16384
	// v2 N=17 R=50.0  T=0s         RD=16 Fc=32768
	// v2 N=18 R=50.0  T=1ms        RD=17 Fc=65536
	// v2 N=19 R=50.0  T=2ms        RD=18 Fc=131072
	// v2 N=20 R=50.0  T=5ms        RD=19 Fc=262144
	// v2 N=21 R=50.0  T=10ms       RD=20 Fc=524288
	// v2 N=22 R=50.0  T=22ms       RD=21 Fc=1048576
	// v2 N=23 R=50.0  T=43ms       RD=22 Fc=2097152
	// v2 N=24 R=50.0  T=84ms       RD=23 Fc=4194304
	// v2 N=25 R=50.0  T=169ms      RD=24 Fc=8388608
	// v2 N=26 R=50.0  T=353ms      RD=25 Fc=16777216
	// v2 N=27 R=50.0  T=678ms      RD=26 Fc=33554432
	// v2 N=28 R=50.0  T=1.377s     RD=27 Fc=67108864
	// v2 N=29 R=50.0  T=2.705s     RD=28 Fc=134217728
	// v2 N=30 R=50.0  T=5.437s     RD=29 Fc=268435456
	// v2 N=31 R=50.0  T=10.856s    RD=30 Fc=536870912
	// v2 N=32 R=50.0  T=21.801s    RD=31 Fc=1073741824
	// v2 N=33 R=50.0  T=44.53s     RD=32 Fc=2147483648
	// v2 N=34 R=50.0  T=1m33.856s  RD=33 Fc=4294967296
	// ...
}
