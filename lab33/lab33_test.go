package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	c, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	fmt.Println(callWithTimeout(c, action))
}

func action() {
	time.Sleep(time.Second * 5)
}

func callWithTimeout(c context.Context, f func()) error {
	var err error
	go func() {
		f()
	}()
loop:
	for {
		select {
		case <-c.Done():
			err = fmt.Errorf("timeout")
			break loop
		default:
			time.Sleep(time.Second)
		}
	}
	return err
}
