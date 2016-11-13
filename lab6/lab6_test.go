package lab6_test

import (
	"fmt"
	"testing"
)

func TestFloatEqual(t *testing.T) {
	var f1 float32 = 1.1111117
	var f2 float32 = 1.1111116
	equal := f1 == f2
	fmt.Printf("%v==%v: %v\n", f1, f2, equal)

}

func TestFloat1(t *testing.T) {
	tt := []uint32{
		0x7FFFFFFE,
		0x7FFFFFFF,
		0x80000000,
		0x80000001,
		0xFF000000,
	}
	for _, u := range tt {
		oki := int32(u)
		f := float64(u)
		fi := int32(f)
		err := oki != fi
		fmt.Printf("x=0x%08X\tu=%10d\toki=%11d\tf=%12.1f\tfi=%11d\t err=%v\n", u, u, oki, f, fi, err)
	}
	//x=0x7FFFFFFE	u=2147483646	oki= 2147483646	f=2147483646.0	fi= 2147483646	 err=false
	//x=0x7FFFFFFF	u=2147483647	oki= 2147483647	f=2147483647.0	fi= 2147483647	 err=false
	//x=0x80000000	u=2147483648	oki=-2147483648	f=2147483648.0	fi=-2147483648	 err=false
	//x=0x80000001	u=2147483649	oki=-2147483647	f=2147483649.0	fi=-2147483648	 err=true
	//x=0xFF000000	u=4278190080	oki=  -16777216	f=4278190080.0	fi=-2147483648	 err=true
}

func TestFloat2(t *testing.T) {
	tt := []float64{
		0x7FFFFFFE,
		0x7FFFFFFF,
		-1,
		-2,
		0x80000000,
		0x80000001,
		0x80000002,
		0x880000002,
		0xFF000000,
		0xFFFFFFFE,
		0xFFFFFFFF,
	}
	for _, f := range tt {
		fi := int32(f)
		u := uint32(f)
		oki := int32(u)
		err := fi != oki
		fmt.Printf("x=0x%08X\tf=%13.1f\tu=%10d\tfi=%11d\toki=%11d\terr=%v\n", u, f, u, fi, oki, err)
	}
	//x=0x7FFFFFFE	f= 2147483646.0	u=2147483646	fi= 2147483646	oki= 2147483646	err=false
	//x=0x7FFFFFFF	f= 2147483647.0	u=2147483647	fi= 2147483647	oki= 2147483647	err=false
	//x=0xFFFFFFFF	f=         -1.0	u=4294967295	fi=         -1	oki=         -1	err=false
	//x=0xFFFFFFFE	f=         -2.0	u=4294967294	fi=         -2	oki=         -2	err=false
	//x=0x80000000	f= 2147483648.0	u=2147483648	fi=-2147483648	oki=-2147483648	err=false
	//x=0x80000001	f= 2147483649.0	u=2147483649	fi=-2147483648	oki=-2147483647	err=true
	//x=0x80000002	f= 2147483650.0	u=2147483650	fi=-2147483648	oki=-2147483646	err=true
	//x=0x80000002	f=36507222018.0	u=2147483650	fi=-2147483648	oki=-2147483646	err=true
	//x=0xFF000000	f= 4278190080.0	u=4278190080	fi=-2147483648	oki=  -16777216	err=true
	//x=0xFFFFFFFE	f= 4294967294.0	u=4294967294	fi=-2147483648	oki=         -2	err=true
	//x=0xFFFFFFFF	f= 4294967295.0	u=4294967295	fi=-2147483648	oki=         -1	err=true
}
