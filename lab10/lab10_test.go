package lab10_test

import (
	"fmt"
	"testing"
)

const (
	use_recursive = false //是否使用递归算法

)

var (
	recursiveCnt = 0
)

//问题：
//一个班有50个人 每个人都有自己固定的座位
//有一天1号疯了 他会随机选择一个座位 后面的人如果自己的座位没被占会坐自己的座位
//如果自己的座位被占了 他也会随机选择一个座位
//
//请编程实现：求最后一个人坐到自己座位的概率
//
//这是递归fuck cpu的又一个例子

func TestR(t *testing.T) {
	if use_recursive {
		N := int(32)
		for i := 1; i <= N; i++ {
			x, maxDepth := R(i, 1)
			fmt.Println("recursive:", i, x, maxDepth, recursiveCnt)
		}
	} else {
		N := int(5000)
		t := []int{}
		for i := 1; i <= N; i++ {
			x, maxDepth := R2(i, &t, 1)
			fmt.Println("no recursive:", i, x, maxDepth, recursiveCnt)
		}
	}
}

func R(n int, depth int) (int, int) { //递归算法 算不出来
	x := int(0)
	maxDepth := 1
	if depth == 1 {
		recursiveCnt = 0
	}
	recursiveCnt++
	for i := int(1); i <= n; i++ {
		xx, md := r(n, i, depth+1)
		x += xx
		if md > maxDepth {
			maxDepth = md
		}
	}
	return x / n, maxDepth
}

func r(n, i int, depth int) (int, int) {
	recursiveCnt++
	switch {
	case i == 1:
		return 100, depth
	case i == n:
		return 0, depth
	default:
		return R(n-i+1, depth+1)
	}
}

func R2(n int, t *[]int, depth int) (int, int) { //非递归算法
	if depth > 3 { //这个也是个递归写法 但是t的存在 可以保证递归深度绝对不会超过3层
		panic(depth)
	}
	if depth == 1 {
		recursiveCnt = 0
	}
	recursiveCnt++
	if len(*t) > n { //已经知道结果的 不用算了 直接返回
		return (*t)[n-1], depth
	}
	x := int(0)
	maxDepth := 1
	for i := int(1); i <= n; i++ {
		tt, md := r2(n, i, t, depth+1)
		x += tt
		if md > maxDepth {
			maxDepth = md
		}
	}
	r := x / n
	*t = append(*t, r)
	return r, maxDepth
}

func r2(n, i int, t *[]int, depth int) (int, int) {
	recursiveCnt++
	switch {
	case i == 1:
		return 100, depth
	case i == n:
		return 0, depth
	default:
		return R2(n-i+1, t, depth+1)
	}
}
