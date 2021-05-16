package main

import (
	"fmt"
	"time"
)

type Version2 struct {
	maxDepth     int
	recursiveCnt int
}

func (v *Version2) R(n, depth int) (percent float64) {
	sum := float64(0)
	for x := int(1); x <= n; x++ {
		s := v.s(n, x, depth)
		if depth > v.maxDepth {
			v.maxDepth = depth
		}
		v.recursiveCnt++
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

func main() {
	NN := 50
	for N := 1; N <= NN; N++ {
		v := &Version2{}
		start := time.Now()
		rate := v.R(N, 1)
		dur := time.Now().Sub(start) / time.Millisecond * time.Millisecond
		fmt.Printf("v2 N=%-2d R=%-5.1f T=%-10s RD=%-2d RN=%d\n", N, rate, dur, v.maxDepth, v.recursiveCnt)
	}
}
