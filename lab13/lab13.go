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
	fmt.Printf("sysNow: %-60s %#v\n", sysNow.String(), timeStruct(sysNow))
	fmt.Printf("myNow : %-60s %#v\n", myNow.String(), timeStruct(myNow))

	d := time.Second * 10
	dest := sysNow.Add(d)
	fmt.Println("dest:", d.String(), dest)
	time.AfterFunc(d, afterFunc)

	t := time.NewTimer(d)
	//modify system time here...
	modifySystime(-3)
	after := <-t.C
	fmt.Printf("after: %-60s %#v\n", after.String(), timeStruct(after))
	fmt.Printf("diffSysNow: %s\n", after.Sub(sysNow)) //not rely to sysytem time
	fmt.Printf("diffMyNow: %s\n", after.Sub(myNow))   //rely to system time

	//sysNow: 2018-11-08 11:30:06.56078 +0800 CST m=+0.006000301           main.timeStruct{wall:0xbef1094fa16cd2e0, ext:6000301, loc:(*time.Location)(0x579300)}
	//myNow : 2018-11-08 11:30:06 +0800 CST                                main.timeStruct{wall:0x0, ext:63677244606, loc:(*time.Location)(0x579300)}
	//dest: 10s 2018-11-08 11:30:16.56078 +0800 CST m=+10.006000301
	//after: 2018-11-08 11:30:09.608378 +0800 CST m=+10.076576301         main.timeStruct{wall:0xbef1095064431c90, ext:10076576301, loc:(*time.Location)(0x579300)}
	//logic diff 10.070576s
	//after: 2018-11-08 11:30:09.608378 +0800 CST m=+10.076576301         main.timeStruct{wall:0xbef1095064431c90, ext:10076576301, loc:(*time.Location)(0x579300)}
	//real diff 3.608378s
	//diffSysNow: 10.070576s
	//diffMyNow: 3.608378s
}

func afterFunc() {
	t := time.Now()
	fmt.Printf("afterf: %-60s %#v\n", t, timeStruct(t)) //10second,but when system time change during sleep, real time is not correct
	fmt.Println("logic diff", t.Sub(start).String())
	myStart := fmtTime(start)
	fmt.Println("real diff", t.Sub(myStart).String())
}
