//去掉符号调试信息
//go build -ldflags "-s"

package main

func main() {
	const debug = true
	if debug {
		println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++this is debug mode============================================================================")
	}
	panic(0)
}
