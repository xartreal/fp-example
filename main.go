package main

import (
	"fmt"
	"strconv"
	"strings"
)

func atoi(in string) int64 {
	q, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return 0
	}
	return q
}

func ftoi(in string) int64 { // fixed->int
	in = strings.ReplaceAll(in, ",", ".")
	if strings.Count(in, ".") > 1 {
		return 0  // 2 or more decimal point
	}
	pos := strings.LastIndex(in, ".")
	if pos == -1 { // no decimal point
		return atoi(in + "00")
	}
	z := len(in) - pos - 1 // digits after decimal point
	if z > 2 { //leave only 2 digits after decimal point
		in = in[0 : pos+3]
	}
	x := atoi(strings.ReplaceAll(in, ".", ""))
	if z == 1 { 
		return x * 10
	}
	return x
}

var signmap = map[bool]string{false: "", true: "-"}

func itof(num int64) string { //int->fixed
	t := strings.TrimPrefix(strconv.FormatInt(num, 10), "-")
	sign := num < 0
	if sign {
		num = -num
	}
	if num < 100 {
		if num < 10 {
			return signmap[sign] + "0,0" + t
		} else {
			return signmap[sign] + "0," + t
		}
	}
	out := t[:len(t)-2] + "," + t[len(t)-2:]
	return signmap[sign] + out
}

//fixed point example

func main() {
	a := ftoi("123.45")
	b := ftoi("345,67")
	c := ftoi("150")
	d := ftoi("151")
	e := ftoi("-200.1")
	r1 := a + b
	r2 := a * 2
	r3 := b - a
	r4 := c / 2
	r5 := -d / 2
	r6 := a - b
	r7 := a - e
	fmt.Printf("r1=%s\n", itof(r1))
	fmt.Printf("r2=%s\n", itof(r2))
	fmt.Printf("r3=%s\n", itof(r3))
	fmt.Printf("r4=%s\n", itof(r4))
	fmt.Printf("r5=%s\n", itof(r5))
	fmt.Printf("---\n")
	fmt.Printf("r6=%s\n", itof(r6))
	fmt.Printf("r7=%s\n", itof(r7))
	fmt.Printf("x1=%s\n", itof(-55))
	fmt.Printf("x2=%s\n", itof(55))
}
