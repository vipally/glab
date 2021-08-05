package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"reflect"
	"testing"
)

func TestBase64(t *testing.T) {
	cs := []*base64.Encoding{
		base64.StdEncoding,
		base64.RawStdEncoding,
		base64.URLEncoding,
		base64.RawURLEncoding,
	}

	for x, c := range cs {
		for i := 1; i < 10; i++ {
			buf := make([]byte, i)
			for j := 0; j < 10; j++ {
				rand.Read(buf)
				s := c.EncodeToString(buf)
				dec, err := c.DecodeString(s)
				if err != nil {
					t.Fatalf("base64 test fail encoding=%d, len=%d data=%x encode=%s err=%s", x, len(buf), buf, s, err)
				}
				if !reflect.DeepEqual(buf, dec) {
					t.Fatalf("base64 test fail encoding=%d, len=%d data=%x encode=%s dec=%x", x, len(buf), buf, s, dec)
				}
				if false {
					fmt.Println(x, i, j, buf, s, dec)
				}
			}
		}
	}
}
