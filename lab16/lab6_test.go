package lab16

import (
	"fmt"
	"testing"
)

type GUID struct {
	Data [4]byte
}

func NewGUID(v uint) GUID {
	var guid GUID
	guid.Set(v)
	return guid
}

func (this *GUID) Set(v uint) {
	this.Data[0] = uint8(v & 0xff)
	this.Data[1] = uint8((v >> 8) & 0xff)
	this.Data[2] = uint8((v >> 16) & 0xff)
	this.Data[3] = uint8((v >> 24) & 0xff)
}

// it's not safe, because it return address of array Data
// and array is value in Go
func (this *GUID) Bytes() []byte {
	return this.Data[:]
}

func getGuid(i uint) []byte {
	var guid GUID
	guid.Set(i)
	return guid.Bytes()
}

// get 1 01000000
// get 2 02000000
// get 3 03000000
// check 1 03000000
// check 2 03000000
// check 3 03000000
func TestArray(t *testing.T) {
	var s [][]byte
	ss := []GUID{NewGUID(1), NewGUID(2), NewGUID(3)}
	for i, x := range ss {
		g := x.Bytes()
		s = append(s, g)
		fmt.Printf("get %d %x\n", i+1, g)
	}
	for i, v := range s {
		fmt.Printf("check %d %x\n", i+1, v)
	}
}
