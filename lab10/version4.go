package main

import (
	"fmt"
	"time"
)

type Version4 struct {
	maxDepth  int
	fnCallCnt int
	known     []float64
}

func (v *Version4) R(n, depth int) (percent float64) {
	if depth > 3 { // assert recursive depth
		panic(depth)
	}

	if depth == 1 {
		v.fnCallCnt = 0
		v.maxDepth = 1
		v.known = v.known[:0]
	}
	if depth > v.maxDepth {
		v.maxDepth = depth
	}

	if n <= len(v.known) { //result is known, use the result directlly
		return v.known[n-1]
	}

	v.fnCallCnt++

	//println(n, depth, len(v.known))
	sum := float64(0)
	//for x := int(1); x <= n; x++ {
	for x := n; x >= 1; x-- { // reverse order to let R(1~n) evaluate order
		s := v.s(n, x, depth)
		sum += s
	}
	r := sum / float64(n)
	v.known = append(v.known, r)

	return r
}

func (v *Version4) s(n, x, depth int) (percent float64) {
	switch {
	case x == 1:
		return 100
	case x == n:
		return 0
	default:
		return v.R(n-x+1, depth+1)
	}
}

func main4() {
	M := 50
	for N := M; N <= M; N++ {
		v := &Version4{}
		start := time.Now()
		rate := v.R(N, 1)
		dur := time.Now().Sub(start) / time.Millisecond * time.Millisecond
		fmt.Printf("v4 N=%-2d R=%-5.1f T=%-10s RD=%-2d Fc=%d\n", N, rate, dur, v.maxDepth, v.fnCallCnt)
	}
	// output:
	// v4 N=50     R=50.0  T=0s         RD=3  Fc=50
	// v4 N=50000  R=50.0  T=10.136s    RD=3  Fc=50000
	// v4 N=100000 R=50.0  T=40.075s    RD=3  Fc=100000
}
