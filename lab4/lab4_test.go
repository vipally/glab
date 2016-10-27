package lab4_test

import (
	"fmt"
	"testing"

	"github.com/vipally/gx/unsafe"
)

func TestStruct(t *testing.T) {
	s := []string{
		"",
		"",
		"hello",
		"hello",
		"hellp",
		fmt.Sprintf("hello"),
		fmt.Sprintf("hello"),
		fmt.Sprintf(""),
		fmt.Sprintf(""),
	}

	for i, v := range s {
		b := unsafe.StringBytes(v)
		b2 := []byte(v)

		if b.Writeable() {
			b[0] = 'x'
		}
		fmt.Println(i, v, unsafe.StringPointer(v), unsafe.BytesPointer(b), unsafe.BytesPointer(b2))
	}

	if true {
		b := []byte{'h', 'e', 'l', 'l', 'o'}
		bs := unsafe.BytesString(b)
		bs2 := string(b)
		fmt.Println(bs2, unsafe.BytesPointer(b), unsafe.StringPointer(string(bs)), unsafe.StringPointer(bs2))
	}

}
