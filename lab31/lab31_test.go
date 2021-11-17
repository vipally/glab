package lab31

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

// 结论：defer之前返回 后置defer函数不会被设置和执行
func testDerfer() (err error) {
	defer func() {
		if err != nil {
			fmt.Println("testDerfer 1", err)
		}
	}()
	return errors.New("error!")
	defer func() {
		if err != nil {
			fmt.Println("testDerfer 2", err)
		}
	}()
	return nil
}

// 结论：同一层:=新建err会复用变量，嵌套代码块中会新建err变量
func testShadow() error {
	var err error
	fmt.Println(&err)
	defer func() {
		if err != nil {
			fmt.Println("testShadow", err, &err)
		}
	}()
	if time.Now().Second()%2 == 0 {
		x, err := 0, errors.New("error inside")
		fmt.Println(&err)
		x = x
		if err != nil {
			return err
		}
	}

	x, err := 0, errors.New("error outside")
	fmt.Println(&err)
	x = x
	if err != nil {
		return err
	}

	return err
}

func TestDefer(t *testing.T) {
	testDerfer()
	testShadow()
}
