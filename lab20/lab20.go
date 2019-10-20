package main

import (
	"fmt"
)

func main() {
	t := "hello 中国！。."
	s := t[:7]
	fmt.Println(len(t))
	for i, v := range t {
		fmt.Println(i, v, string(v), []byte(string(v)))
	}
	for i := 0; i < len(t); i++ {
		fmt.Println(i, t[i])
	}

	fmt.Println(len(s))
	for i, v := range s {
		fmt.Println(i, v, string(v), []byte(string(v)))
	}
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
}
