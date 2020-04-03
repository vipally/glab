package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	strings := []string{"", "", ""}
	flag.StringVar(&strings[0], "a", "", "a")
	flag.StringVar(&strings[1], "b", "", "b")
	flag.StringVar(&strings[2], "c", "", "c")
	flag.Parse()
	fmt.Printf("args: %#v\n", os.Args)
	fmt.Printf("got: %#v\n", strings)
	//lab21 -a="hello -world" -b="--help" -c="=abc="
}
