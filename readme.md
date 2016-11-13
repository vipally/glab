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


	
