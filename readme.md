# glab
some test code of golang

## lab1: reflect

## lab2: Recursive
	go test -test.bench=".*"
	N = 45
	testing: warning: no tests to run
	Benchmark_Recursive-4                  1        10018573100 ns/op
	Benchmark_Loop-4                2000000000               0.00 ns/op
	Benchmark_Stack-4               2000000000               0.00 ns/op
	PASS
	ok      github.com/vipally/glab/lab2    10.154s
	
	结论：
	本测试求Fibonacci(45) 递归算法已经无法忍受 当N再增加到50以上的时候 递归算法已经无法算出结果
	而采用循环和stack的算法 N=10亿 仍然毫无压力
	
## lab2: Golang deep

## lab3: size test

## lab4: string buffer test

## lab5: functor test

## lab6: float test

## lab7: interface inherit test

## lab8: interface nil test
	含有匿名成员的结构体 可以由父对象默认实现某个接口 达到实现某接口默认行为
	子对象可以实现该函数 覆盖父对象的默认行为

## lab9: go parse

## lab10: disaster of recursive

## lab11: constant branch optimization 常量条件编译优化

## lab12: confused interface

## lab13: timer test during system time rechange

## lab14: channel and mutex test on 1WnR case

## lab15: random for test dice

## lab16: GUID:array.slice error when take address of range value
	BUG:对range变量x进行取地址将会是一场灾难 因为go是值拷贝 并且range变量会在遍历期间重用
	如果对该变量取地址 那么存下来的东西就只剩集合里面最后一个元素了
	这个问题可以跟Go提一个issue，禁止对range变量取地址

## lab17: Go dll test

## lab18: range slice & range map speed test
	结论：对同样大小的slice和map执行range，对map的遍历慢10倍

## lab19 goroutine test
## lab20 string test
 对string range，得到的是rune字符序列，index不一定连续
 对string取下标 得到的是对应[]byte的字节


	
