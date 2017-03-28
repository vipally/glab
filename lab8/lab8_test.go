package lab8_test

import (
	"fmt"
	"testing"
)

type Iter interface {
	Do()
}

type A struct {
}

func (this *A) Do() {
	fmt.Println("A.Do")
}

func TestInterface(t *testing.T) {
	var i Iter
	var i2 Iter = (*A)(nil)
	fmt.Printf("i=%#v i2=%#v equal=%v\n", i, i2, i == i2)

	//i=<nil> i2=(*lab8_test.A)(nil) equal=false
}
