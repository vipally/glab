package lab3_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type empty struct{}
type withEmpty struct {
	a uint
	c empty
	d int
	b uint16
}

func TestStruct(t *testing.T) {
	fmt.Println("empty", unsafe.Sizeof(empty{})) //空
	fmt.Println("byte", unsafe.Sizeof(byte(0)))
	fmt.Println("int", unsafe.Sizeof(int(0)))
	fmt.Println("withEmpty", unsafe.Sizeof(withEmpty{}))
	//var we withEmpty

	//empty 0
	//byte 1
	//int 8
	//withEmpty 24
	//结论 空结构体size=0 是不是说不需要成员只需要类型信息的地方用empty会代价更小？
	//最后一个字段即使不够8字节 也会补齐8字节？
}

func TestSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := a[0:1]
	c := a[1:2]
	fmt.Println(&a[0], &b[0], &c[0])
}
