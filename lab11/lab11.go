//去掉符号调试信息
//go build -ldflags "-s"

package main

func main() {
	const debug = false

	//这里条件是编译期常量 编译器不能优化吗 如果条件一直是false 就忽略这段代码 不用生成 这不就是C的条件编译了
	//ref:c++11 constexp
	if debug {
		println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++this is debug mode============================================================================")
	}
	panic(0)
}
