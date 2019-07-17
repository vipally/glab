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

## lab10: 递归的危害

## lab11: constant branch optimization 常量条件编译优化

## lab12: confused interface

## lab18: range slice & range map speed test
	结论：对同样大小的slice和map执行range，对map的遍历慢10倍


	
