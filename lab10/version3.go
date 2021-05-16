package main

import (
	"fmt"
	"time"
)

type Version3 struct {
	maxDepth  int
	fnCallCnt int
	known     []float64
}

func (v *Version3) R(n, depth int) (percent float64) {
	if depth == 1 {
		v.fnCallCnt = 0
		v.maxDepth = 1
		v.known = v.known[:0]
	}
	if len(v.known) >= n { //result is known, use the result
		return v.known[n-1]
	}
	//println(n, depth, len(v.known))
	if depth > v.maxDepth {
		v.maxDepth = depth
	}
	v.fnCallCnt++

	sum := float64(0)
	for x := int(1); x <= n; x++ {
		s := v.s(n, x, depth)
		sum += s
	}
	r := sum / float64(n)
	v.known = append(v.known, r)
	return r
}

func (v *Version3) s(n, x, depth int) (percent float64) {
	switch {
	case x == 1:
		return 100
	case x == n:
		return 0
	default:
		return v.R(n-x+1, depth+1)
	}
}

func main3() {
	M := 50
	for N := M; N <= M; N++ {
		v := &Version3{}
		start := time.Now()
		rate := v.R(N, 1)
		dur := time.Now().Sub(start) / time.Millisecond * time.Millisecond
		fmt.Printf("v3 N=%-2d R=%-5.1f T=%-10s RD=%-2d Fc=%d\n", N, rate, dur, v.maxDepth, v.fnCallCnt)
	}
	// output:
	// v3 N=50     R=50.0  T=0s         RD=49 Fc=49
	// v3 N=50000  R=50.0  T=10.254s    RD=49999 Fc=49999
	// v3 N=100000 R=50.0  T=39.489s    RD=99999 Fc=99999
}
