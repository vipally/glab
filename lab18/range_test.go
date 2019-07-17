package lab18

import (
	"fmt"
	"testing"
)

// testSize=1000000 testNum=30
// goos: windows
// goarch: amd64
// pkg: github.com/vipally/glab/lab18
// BenchmarkSliceRange-6   	2000000000	         0.02 ns/op
// BenchmarkMapRange-6     	2000000000	         0.21 ns/op

// 结论：很明显对map做range比对slice range慢太多，所以没有做查询的需求 尽量用slice解决问题

var mp map[int]string
var sl []string

const testNum = 30

func init() {
	max := 1000000
	mp = make(map[int]string)
	sl = make([]string, max)
	for i := 1; i <= max; i++ {
		s := fmt.Sprint(i)
		mp[i] = s
		sl[i-1] = s
	}
	fmt.Printf("testSize=%d testNum=%d\n", max, testNum)
}

func BenchmarkSliceRange(b *testing.B) {
	for i := 0; i < testNum; i++ {
		for _, v := range sl {
			v = v
		}
	}
}

func BenchmarkMapRange(b *testing.B) {
	for i := 0; i < testNum; i++ {
		for _, v := range mp {
			v = v
		}
	}
}
