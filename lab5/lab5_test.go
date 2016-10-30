package lab5_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type CmpFunc func(left, right int) bool

//create cmp object by name
func GetCmpFunc(cmpName string) (r CmpFunc) {
	switch cmpName {
	case "": //default Lesser
		fallthrough
	case "Lesser":
		r = Less
	case "Greater":
		r = Great
	default: //unsupport name
		panic(cmpName)
	}
	return
}

//Lesser
func Less(left, right int) (ok bool) {
	ok = left < right
	return
}

//Greater
func Great(left, right int) (ok bool) {
	ok = right > left
	return
}

type Comparer interface {
	F(left, right int) bool
}

//create cmp object by name
func CreateComparer(cmpName string) (r Comparer) {
	switch cmpName {
	case "": //default Lesser
		fallthrough
	case "Lesser":
		r = Lesser{}
	case "Greater":
		r = Greater{}
	default: //unsupport name
		panic(cmpName)
	}
	return
}

//Lesser
type Lesser struct{}

func (this Lesser) F(left, right int) (ok bool) {
	ok = left < right
	return
}

//Greater
type Greater struct{}

func (this Greater) F(left, right int) (ok bool) {
	ok = right < left
	return
}

type CmpObj byte

const (
	CMP_LESS CmpObj = iota
	CMP_GREAT
)

func (me CmpObj) F(left, right int) (ok bool) {
	switch me {
	case CMP_LESS:
		ok = Less(left, right)
	case CMP_GREAT:
		ok = Great(left, right)
	default:
		panic(me)
	}
	return
}

var (
	cmp1  = CreateComparer("Lesser")
	cmp2  = GetCmpFunc("Lesser")
	cmp3  = CMP_LESS
	start = 0
	N     = 100000000
)

func TestSize(t *testing.T) {
	fmt.Println("Interface", unsafe.Sizeof(cmp1))
	fmt.Println("Func", unsafe.Sizeof(cmp2))
	fmt.Println("Obj", unsafe.Sizeof(cmp3))
}

func Benchmark_Interface(b *testing.B) {
	for i := 0; i < N; i++ {
		cmp1.F(1, 2)
	}
}
func Benchmark_Func(b *testing.B) {
	for i := 0; i < N; i++ {
		cmp2(1, 2)
	}
}
func Benchmark_Obj(b *testing.B) {
	for i := 0; i < N; i++ {
		cmp3.F(1, 2)
	}
}

//Interface 16
//Func 8
//Obj 1
//Benchmark_Interface-4   	1000000000	         0.55 ns/op
//Benchmark_Func-4        	2000000000	         0.19 ns/op
//Benchmark_Obj-4         	2000000000	         0.24 ns/op
//
//结论：
//用interface实现多态，会占用两个指针(16字节空间) 执行效率上 大概慢一倍
//使用函数指针 占用1个指针(8字节) 执行效率最高
//使用转调对象 只需要1个字节 执行效率跟函数指针差不多
