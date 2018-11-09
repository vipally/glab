package main

import (
	"fmt"
	"os/exec"
	"time"
)

var start time.Time

type timeStruct time.Time

func fmtTime(t time.Time) time.Time {
	fmt_ := "2006-01-02 15:04:05"
	fmtT, _ := time.ParseInLocation(fmt_, t.Format(fmt_), time.Local)
	return fmtT
}
func modifySystime(dur int) {
	fmt.Println("modifySystime", dur)
	now := time.Now()
	dest := now.Add(time.Duration(dur) * time.Second)
	para := fmt.Sprintf("%02d:%02d:%02d", dest.Hour(), dest.Minute(), dest.Second())
	p := exec.Command("cmd", "/C", "time", para)
	err := p.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	sysNow := time.Now()
	start = sysNow
	myNow := fmtTime(sysNow)
	fmt.Printf("sysNow: %-60s %d %#v\n", sysNow.String(), sysNow.Unix(), timeStruct(sysNow))
	fmt.Printf("myNow : %-60s %d %#v\n", myNow.String(), myNow.Unix(), timeStruct(myNow))

	d := time.Second * 10
	dest := sysNow.Add(d)
	fmt.Printf("dest : %s %-60s %d %#v\n", d.String(), dest.String(), dest.Unix(), timeStruct(dest))
	time.AfterFunc(d, afterFunc)

	t := time.NewTimer(d)
	//modify system time here...
	modify := -3
	modifySystime(modify)
	after := <-t.C
	fmt.Printf("after: %-60s %d %#v\n", after.String(), after.Unix(), timeStruct(after))
	fmt.Printf("diffSysNow: %s\n", after.Sub(sysNow)) //not rely to sysytem time
	fmt.Printf("diffMyNow: %s\n", after.Sub(myNow))   //rely to system time

	modifySystime(-modify)

	//sysNow: 2018-11-09 23:52:48.7326954 +0800 CST m=+0.015600101         1541778768 main.timeStruct{wall:0xbef189342bac0b68, ext:15600101, loc:(*time.Location)(0x579c20)}
	//myNow : 2018-11-09 23:52:48 +0800 CST                                1541778768 main.timeStruct{wall:0x0, ext:63677375568, loc:(*time.Location)(0x579c20)}
	//dest:10s 2018-11-09 23:52:58.7326954 +0800 CST m=+10.015600101       1541778778 main.timeStruct{wall:0xbef18936abac0b68, ext:10015600101, loc:(*time.Location)(0x579c20)}
	//modifySystime -3
	//after: 2018-11-09 23:52:54.9771123 +0800 CST m=+10.047314601         1541778774 main.timeStruct{wall:0xbef18935ba3d8cec, ext:10047314601, loc:(*time.Location)(0x579c20)}
	//diffSysNow: 10.0317145s
	//diffMyNow: 6.9771123s
	//modifySystime 3
	//afterf: 2018-11-09 23:52:54.9771123 +0800 CST m=+10.047314601        1541778774 main.timeStruct{wall:0xbef18935ba3d8cec, ext:10047314601, loc:(*time.Location)(0x579c20)}
	//logic diff 10.0317145s
	//real diff 6.9771123s
}

func afterFunc() {
	t := time.Now()
	fmt.Printf("afterf: %-60s %d %#v\n", t, t.Unix(), timeStruct(t)) //10second,but when system time change during sleep, real time is not correct
	fmt.Println("logic diff", t.Sub(start).String())
	myStart := fmtTime(start)
	fmt.Println("real diff", t.Sub(myStart).String())
}
