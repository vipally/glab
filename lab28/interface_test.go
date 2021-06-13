package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

/*
type MyInt int

// It's ok to define method for normal type
func (m MyInt) Show() {
	println(m)
}

type MyInterface interface{}

// invalid receiver type MyInterface (MyInterface is an interface type)
func (m MyInterface) Show() {
	fmt.Println(m)
}

func (m *MyInterface) UnmarshalJSON(data []byte) error {
	*m = append([]byte(nil), data...)
	return nil
}
*/

func TestUniversalUnmarshal(t *testing.T) {
	var d struct {
		D interface{}
	}
	var txt = `
{
	"D":{
	  "a": 1,
	  "b": "foo"
	}
}
`
	json.Unmarshal([]byte(txt), &d)
	fmt.Printf("unmarshal txt: %s\nGo data:%#v\n", txt, d)
	// output:
	// unmarshal txt:
	// {
	// "D":{
	//   "a": 1,
	//   "b": "foo"
	// }
	// }
	// Go data:struct { D interface {} }{D:map[string]interface {}{"a":1, "b":"foo"}}
}

func TestUniversalUnmarshal2(t *testing.T) {
	var d interface{}
	var txt = `
{
  "a": 1,
  "b": "foo"
}
`
	err := json.Unmarshal([]byte(txt), d)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("unmarshal txt: %s\nGo data:%#v, err=%v\n", txt, d, err)
	// output:
	// json: Unmarshal(nil)
}

func TestInterfaceReflect(t *testing.T) {
	var d interface{}
	Unmarshal(nil, d)
	Unmarshal(nil, &d)
	// output:
	// Unmarshal v=<nil> rt=<nil> rv=<invalid reflect.Value>
	// Unmarshal v=(*interface {})(0xc00003a740) rt=*interface {} rv=0xc00003a740
	//  rt=ptr rv=0xc00003a740
	//    rtreal=interface rv=0xc00003a740
}

func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)

	fmt.Printf("Unmarshal v=%#v rt=%v rv=%v\n", v, rt, rv)
	if rt != nil {
		fmt.Printf("  rt=%v rv=%v\n", rt.Kind(), rv)
		if rt.Kind() == reflect.Ptr {
			rtReal := rt.Elem()
			fmt.Printf("    rtreal=%v rv=%v\n", rtReal.Kind(), rv)
		}
	}
	return nil
}
