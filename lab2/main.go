package main

import (
	"time"
)

var cnt int

//请评估以下函数的执行时间
func fibonacci(n uint64) uint64 {
	cnt++
	if n > 2 {
		return fibonacci(n-1) + fibonacci(n-2)
	}
	return 1
}
func fact(n int) int {
	v := 1
	for i := n; i > 1; i-- {
		v *= i
	}
	return v
}
func main() {
	start := time.Now()
	f60 := fibonacci(9)
	dur := time.Now().Sub(start)
	println("f60=", f60, "cost=", dur.String(), cnt)

}

/*
计算递归函数的调用次数
1 c=1
2 c=1
3 c(2)+c(1)+1=3
4 c(3)+c(2)+1=5
5 c(4)+c(3)+1=9
6 c(5)+c(4)+1=15
7 c(6)+c(5)+1=25
8 c(7)+c(6)+1=41=21+20
9 67=34+33
10 109=55+54
*/
