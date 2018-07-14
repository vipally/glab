//去掉符号调试信息
//go build -ldflags "-s"

//test "constant branch optimization"
//constexpr

package main

func main() {
	const (
		debug = false
		lang  = "CHN"
		os    = "windows"
	)

	//这里条件是编译期常量 编译器不能优化吗 如果条件一直是false 就忽略这段代码 不用生成 这不就是C的条件编译了
	//ref:c++11 constexp
	// equal to:
	// println("mode=release")
	if debug {
		println("mode=debug")
	} else {
		println("mode=release")
	}

	//由于lang是常量 一下语句可以去掉switch
	// equal to:
	// println("os=windows")
	switch os {
	case "windows":
		println("os=windows")
	case "linux":
		println("os=linux")
	}

	//由于lang是常量 一下语句可以去掉switch
	// equal to:
	// println("language=Chinese")
	switch lang {
	case "CHN":
		println("language=Chinese")
	case "ENG":
		println("language=English")
	case "JP":
		println("language=Japanese")
	case "VIET":
		println("language=Vietnamese")
	}

	//panic(0)
}
