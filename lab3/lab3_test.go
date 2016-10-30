package lab3_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type empty struct{}

func (me empty) Hello() string { return "" }

type fun func(int, int) int

type withEmpty struct {
	//a uint16
	//c empty
	//b uint16
	//d, e, f, g, h byte
	i I
	f fun
}

type I interface {
	Hello() string
}

func TestStruct(t *testing.T) {
	fmt.Println("empty", unsafe.Sizeof(empty{})) //空
	fmt.Println("byte", unsafe.Sizeof(byte(0)))
	fmt.Println("int", unsafe.Sizeof(int(0)))

	var we withEmpty
	var i I
	fmt.Println("withEmpty", unsafe.Alignof(we), unsafe.Sizeof(we))
	fmt.Println("interface", unsafe.Alignof(i), unsafe.Sizeof(i))

	//empty 0
	//byte 1
	//int 8
	//withEmpty 24
	//结论 空结构体size=0 是不是说不需要成员只需要类型信息的地方用empty会代价更小？
	//最后一个字段即使不够8字节 也会补齐8字节？
	//结构体的Alignof是成员中最大的那位 最后结构体长度是Alignof的整数倍
	//字节对齐最大为8 最小为1
	//empty字段不占用存储空间
	//就算内存空洞在最后，生成array的时候，也可以保证所有对象都是遵循对齐规则的
	//iterface占用16字节 (pType，pValue两个指针)
	//函数占用8个字节 (一个指针)
}

func TestSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := a[0:1]
	c := a[1:2]
	fmt.Println(&a[0], &b[0], &c[0])
}
