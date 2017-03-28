package lab7_test

import (
	"fmt"
	"testing"
)

type Iter interface {
	Do1()
	Do2()
}
type father struct {
}

func (this *father) Do1() { fmt.Println("father Do1") }
func (this *father) Do2() {
	fmt.Println("father Do2")
	this.Do1()
}

type child struct {
	father
}

func (this *child) Do1() {
	//this.father.Do1()
	fmt.Println("child Do1")
}

func TestInterface(t *testing.T) {
	var i Iter = new(child)
	//i.Do1()
	i.Do2()

	//father Do1
	//child Do1
	//father Do2
}
