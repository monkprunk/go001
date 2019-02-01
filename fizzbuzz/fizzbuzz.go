package main

import (
	"fmt"
	"strconv"
)

func Say(i int) string {
	var s string
	if i%3 == 0 {
		s += "Fizz"
	}
	if i%5 == 0 {
		s += "Buzz"
	}
	if s == "" {
		s = strconv.Itoa(i)
	}
	return s
}

func main() {
	fmt.Printf("%#v\n", Say(1))
	fmt.Printf("%#v\n", Say(2))
	fmt.Printf("%#v\n", Say(3))
	fmt.Printf("%#v\n", Say(4))
	fmt.Printf("%#v\n", Say(5))
	fmt.Printf("%#v\n", Say(6))
	fmt.Printf("%#v\n", Say(7))
	fmt.Printf("%#v\n", Say(8))
	fmt.Printf("%#v\n", Say(9))
	fmt.Printf("%#v\n", Say(10))
	fmt.Printf("%#v\n", Say(11))
	fmt.Printf("%#v\n", Say(12))
	fmt.Printf("%#v\n", Say(13))
	fmt.Printf("%#v\n", Say(14))
	fmt.Printf("%#v\n", Say(15))

}
