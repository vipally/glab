package lab15

import (
	"fmt"
	"testing"
)

const (
	randCount = 1000000
)

// little   p=0 count=486801/1000000 rate=48.68%
// big      p=0 count=485346/1000000 rate=48.53%
// baozi    p=1 count=4835/1000000 rate=0.48%
// baozi    p=2 count=4602/1000000 rate=0.46%
// baozi    p=3 count=4444/1000000 rate=0.44%
// baozi    p=4 count=4430/1000000 rate=0.44%
// baozi    p=5 count=4662/1000000 rate=0.46%
// baozi    p=6 count=4880/1000000 rate=0.48%
// point    p=3 count=4835/1000000 rate=0.48%
// point    p=4 count=14069/1000000 rate=1.40%
// point    p=5 count=28123/1000000 rate=2.81%
// point    p=6 count=46551/1000000 rate=4.65%
// point    p=7 count=69324/1000000 rate=6.93%
// point    p=8 count=97853/1000000 rate=9.78%
// point    p=9 count=115485/1000000 rate=11.54%
// point    p=10 count=124442/1000000 rate=12.44%
// point    p=11 count=124379/1000000 rate=12.43%
// point    p=12 count=115953/1000000 rate=11.59%
// point    p=13 count=96746/1000000 rate=9.67%
// point    p=14 count=69012/1000000 rate=6.90%
// point    p=15 count=46388/1000000 rate=4.63%
// point    p=16 count=27798/1000000 rate=2.77%
// point    p=17 count=14162/1000000 rate=1.41%
// point    p=18 count=4880/1000000 rate=0.48%
func TestRand(t *testing.T) {
	var (
		little = 0
		big    = 0
		baozi  [6]int
		point  [16]int
	)
	for i := 0; i < randCount; i++ {
		_, p, bz := randDice()
		if !bz {
			if p <= 10 {
				little++
			} else {
				big++
			}
		} else {
			baozi[(p/3)-1]++
		}
		point[p-3]++
	}
	showResult("little", 0, little)
	showResult("big", 0, big)
	for i, v := range baozi {
		showResult("baozi", 1+i, v)
	}
	for i, v := range point {
		showResult("point", 3+i, v)
	}
}

func showResult(name string, point int, count int) {
	rate := count * 10000 / randCount
	fmt.Printf("%-8s p=%d count=%d/%d rate=%d.%d%%\n", name, point, count, randCount, rate/100, rate%100)
}

// 押注方案 100(大小) x 2 + 0(豹子)*6
// 01 times=058 WinCoin=-200 WinScore=11800 ScorePerCoin=59.00
// 02 times=149 WinCoin=-200 WinScore=30000 ScorePerCoin=150.00
// 03 times=018 WinCoin=-200 WinScore=03800 ScorePerCoin=19.00
// 04 times=107 WinCoin=-200 WinScore=21600 ScorePerCoin=108.00
// 05 times=076 WinCoin=-200 WinScore=15400 ScorePerCoin=77.00
// 06 times=042 WinCoin=-200 WinScore=08600 ScorePerCoin=43.00
// 07 times=026 WinCoin=-200 WinScore=05400 ScorePerCoin=27.00
// 08 times=001 WinCoin=-200 WinScore=00400 ScorePerCoin=2.00
// 09 times=012 WinCoin=-200 WinScore=02600 ScorePerCoin=13.00
// 10 times=066 WinCoin=-200 WinScore=13400 ScorePerCoin=67.00
// totalBet=565 totalWinCoin=-2000 totalWinScore=113000, totalScorePerCoin=56.50
// 押注方案 97(大小) x 2 + 1(豹子)*6
// 01 times=007 WinCoin=-091 WinScore=01600 ScorePerCoin=17.58
// 02 times=062 WinCoin=-421 WinScore=12600 ScorePerCoin=29.93
// 03 times=104 WinCoin=-673 WinScore=21000 ScorePerCoin=31.20
// 04 times=069 WinCoin=-463 WinScore=14000 ScorePerCoin=30.24
// 05 times=075 WinCoin=-499 WinScore=15200 ScorePerCoin=30.46
// 06 times=050 WinCoin=-349 WinScore=10200 ScorePerCoin=29.23
// 07 times=050 WinCoin=-349 WinScore=10200 ScorePerCoin=29.23
// 08 times=000 WinCoin=-049 WinScore=00200 ScorePerCoin=4.08
// 09 times=004 WinCoin=-073 WinScore=01000 ScorePerCoin=13.70
// 10 times=065 WinCoin=-439 WinScore=13200 ScorePerCoin=30.07
// totalBet=496 totalWinCoin=-3406 totalWinScore=99200, totalScorePerCoin=29.13
// 押注方案 94(大小) x 2 + 2(豹子)*6
// 01 times=008 WinCoin=0006 WinScore=01800 ScorePerCoin=-
// 02 times=021 WinCoin=-150 WinScore=04400 ScorePerCoin=29.33
// 03 times=014 WinCoin=-066 WinScore=03000 ScorePerCoin=45.45
// 04 times=039 WinCoin=-366 WinScore=08000 ScorePerCoin=21.86
// 05 times=015 WinCoin=-078 WinScore=03200 ScorePerCoin=41.03
// 06 times=004 WinCoin=0054 WinScore=01000 ScorePerCoin=-
// 07 times=023 WinCoin=-174 WinScore=04800 ScorePerCoin=27.59
// 08 times=005 WinCoin=0042 WinScore=01200 ScorePerCoin=-
// 09 times=015 WinCoin=-078 WinScore=03200 ScorePerCoin=41.03
// 10 times=015 WinCoin=-078 WinScore=03200 ScorePerCoin=41.03
// totalBet=169 totalWinCoin=-888 totalWinScore=33800, totalScorePerCoin=38.06
// 押注方案 91(大小) x 2 + 3(豹子)*6
// 01 times=062 WinCoin=-863 WinScore=12600 ScorePerCoin=14.60
// 02 times=018 WinCoin=-071 WinScore=03800 ScorePerCoin=53.52
// 03 times=071 WinCoin=-1025 WinScore=14400 ScorePerCoin=14.05
// 04 times=027 WinCoin=-233 WinScore=05600 ScorePerCoin=24.03
// 05 times=017 WinCoin=-053 WinScore=03600 ScorePerCoin=67.92
// 06 times=029 WinCoin=-269 WinScore=06000 ScorePerCoin=22.30
// 07 times=004 WinCoin=0181 WinScore=01000 ScorePerCoin=-
// 08 times=007 WinCoin=0127 WinScore=01600 ScorePerCoin=-
// 09 times=016 WinCoin=-035 WinScore=03400 ScorePerCoin=97.14
// 10 times=003 WinCoin=0199 WinScore=00800 ScorePerCoin=-
// totalBet=264 totalWinCoin=-2042 totalWinScore=52800, totalScorePerCoin=25.86
func TestScore(t *testing.T) {
	const (
		baozi     = 0
		littleBig = 100 - baozi*3
	)

	fmt.Printf("押注方案 %d(大小) x 2 + %d(豹子)*6\n", littleBig, baozi)

	totalBet := 0
	totalWinCoin := 0
	totalWinScore := 0
	for i := 0; i < 10; i++ {
		n := getWinTimes()
		WinCoin := 150*baozi - n*6*baozi - littleBig*2 - baozi*5
		WinScore := 200 * (n + 1)
		totalBet += (n + 1)
		totalWinCoin += WinCoin
		totalWinScore += WinScore
		ScorePerCoin := "-"
		if WinCoin < 0 {
			ScorePerCoin = fmt.Sprintf("%.2f", float64(WinScore)/float64(-WinCoin))
		}
		fmt.Printf("%02d times=%03d WinCoin=%04d WinScore=%05d ScorePerCoin=%s\n", i+1, n, WinCoin, WinScore, ScorePerCoin)
	}
	totalScorePerCoin := "-"
	if totalWinCoin < 0 {
		totalScorePerCoin = fmt.Sprintf("%.2f", float64(totalWinScore)/float64(-totalWinCoin))
	}
	fmt.Printf("totalBet=%d totalWinCoin=%d totalWinScore=%d, totalScorePerCoin=%s\n", totalBet, totalWinCoin, totalWinScore, totalScorePerCoin)
}

func getWinTimes() int {
	count := 0
	for {
		_, _, bz := randDice()
		if bz {
			return count
		}
		count++
	}
}
