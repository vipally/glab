package main

import (
	"fmt"
	"testing"
)

//Recursive
const N uint64 = 47

func init() {
	fmt.Println("N =", N)
}

func Benchmark_Recursive(b *testing.B) {
	fibonacci(N)
}
func Benchmark_Loop(b *testing.B) {
	fibonacci2(N)
}
func Benchmark_Stack(b *testing.B) {
	fibonacci3(N)
}

func fibonacci2(n uint64) uint64 {
	f0, f1 := uint64(0), uint64(1)
	for i := uint64(2); i <= n; i++ {
		f0, f1 = f1, f0+f1
	}
	return f1
}

type Stack []uint64

func (this *Stack) push(v uint64) {
	*this = append(*this, v)
}
func (this *Stack) pop() (v uint64, ok bool) {
	l := len(*this)
	if l > 0 {
		v = (*this)[l-1]
		*this = (*this)[:l-1]
		ok = true
	}
	return
}

func fibonacci3(n uint64) uint64 {
	var stack Stack
	for j := n - 1; j >= 2; j-- {
		stack.push(j)
	}
	f0, f1 := uint64(0), uint64(1)
	for _, ok := stack.pop(); ok; _, ok = stack.pop() {
		f0, f1 = f1, f0+f1
	}
	return f1
}

func fibonacci(n uint64) uint64 {
	x := n
	if n > 1 {
		x = fibonacci(n-1) + fibonacci(n-2)
	}
	return x
}
