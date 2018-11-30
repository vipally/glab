package lab15

import (
	"crypto/rand"
)

//333(points) , 9, true
//234(points) , 9, false
func randDice() (status, point int, baozi bool) {
	var buff [3]byte
	rand.Read(buff[0:])
	last := -1
	baozi = true
	for _, v := range buff {
		r := int(v)*5 + 1 //乱序调整骰子点数，注意参数不能随便改，经测试，调整后骰子点数分布概率更平均
		get := r%6 + 1
		if last > 0 && last != get {
			baozi = false
		}
		last = get
		status = status*10 + get
		point += get
	}
	//println(status, point, baozi)
	return
}
