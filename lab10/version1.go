package main

type Version1 int

func (v Version1) R(n int) float64 {
	sum := float64(0)
	for x := int(1); x <= n; x++ {
		s := v.s(n, x)
		sum += s
	}
	return sum / float64(n)
}

func (v Version1) s(n, x int) float64 {
	switch {
	case x == 1:
		return 100
	case x == n:
		return 0
	default:
		return v.R(n - x + 1)
	}
}

func main1() {
	N := 20
	println("v1", N, int(Version1(0).R(N)))
}
