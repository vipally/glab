package lab31

import (
	"errors"
	"fmt"
	"testing"
)

func testDerfer() (err error) {
	defer func() {
		if err != nil {
			fmt.Println("defer1", err)
		}
	}()
	return errors.New("error!")
	defer func() {
		if err != nil {
			fmt.Println("defer2", err)
		}
	}()
	return nil
}

func TestDefer(t *testing.T) {
	testDerfer()
}
