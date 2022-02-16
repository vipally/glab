package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"testing"
	"time"
)

const testN = 1000000

func TestMapSlice(t *testing.T) {
	testOne("Slice", testSlice)
	testOne("  Map", testMap)
}

func testMap() error {
	var data = `
{
	"a":123,
	"b":123,
	"c":123,
	"d":123,
	"e":123,
	"f":123,
	"g":123,
	"h":123,
	"i":123,
	"j":123,
	"k":123,
	"l":123,
	"m":123,
	"n":123,
	"o":123,
	"p":123,
	"q":123
}
`
	b := []byte(data)

	var set [testN]map[string]int
	for i := 0; i < len(set); i++ {
		x := make(map[string]int)
		if err := json.Unmarshal(b, &x); err != nil {
			return err
		}

		// range time
		s := 0
		for _, v := range x {
			s += v
		}

		set[i] = x
	}
	return nil
}

func testSlice() error {
	var data = `
[
	{"n":"a","v":123},
	{"n":"b","v":123},
	{"n":"c","v":123},
	{"n":"d","v":123},
	{"n":"e","v":123},
	{"n":"f","v":123},
	{"n":"g","v":123},
	{"n":"h","v":123},
	{"n":"i","v":123},
	{"n":"j","v":123},
	{"n":"k","v":123},
	{"n":"l","v":123},
	{"n":"m","v":123},
	{"n":"n","v":123},
	{"n":"o","v":123},
	{"n":"p","v":123},
	{"n":"q","v":123}
]
`
	b := []byte(data)
	type Field struct {
		Name  string `json:"n"`
		Value int    `json:"v"`
	}

	var set [testN][]Field
	for i := 0; i < len(set); i++ {
		x := make([]Field, 0)
		if err := json.Unmarshal(b, &x); err != nil {
			return err
		}

		// range time
		s := 0
		for _, v := range x {
			s += v.Value
		}

		set[i] = x
	}
	return nil
}

func testOne(name string, f func() error) {
	runtime.GC()
	start := time.Now()
	if err := f(); err != nil {
		panic(err)
	}
	dur := time.Now().Sub(start)
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	fmt.Printf("test%s N=%d dur=%s mem=%+v\n", name, testN, dur, stats)
	runtime.GC()
}
