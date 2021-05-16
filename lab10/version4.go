package main

import (
	"fmt"
	"time"
)

type Version4 struct {
	maxDepth  int
	fnCallCnt int
	known     []float64
	stk       Stack
}

type Stack []int

func (s *Stack) Push(n int) {
	*s = append(*s, n)
}
func (s *Stack) Pop() int {
	if sz := len(*s); sz > 0 {
		n := (*s)[sz-1]
		*s = (*s)[:sz-1]
		return n
	}
	return -1
}
func (s Stack) Size() int {
	return len(s)
}
func (s Stack) Empty() bool {
	return s.Size() == 0
}

func (v *Version4) R(n, depth int) (percent float64) {
	if depth > 3 { // limit depth
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
	switch {
	case n <= len(v.known): //result is known, use the result
		return v.known[n-1]
	case depth == 1 && n > 2:
		for i := len(v.known) + 1; i < n; i++ { // reverse the call direction like stack
			v.R(i, 2)
		}
	default:
		// do nothing
	}

	v.fnCallCnt++

	//println(n, depth, len(v.known), v.stk.Size())
	sum := float64(0)
	for x := int(1); x <= n; x++ {
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
	NN := 100000
	for N := NN; N <= NN; N++ {
		v := &Version4{}
		start := time.Now()
		rate := v.R(N, 1)
		dur := time.Now().Sub(start) / time.Millisecond * time.Millisecond
		fmt.Printf("v4 N=%-2d R=%-5.1f T=%-10s RD=%-2d Fc=%d\n", N, rate, dur, v.maxDepth, v.fnCallCnt)
	}
	// output:

}
