package fizzbuzz

import "strconv"

func Say(i int) string {
	var s string

	if (i % 3) == 0 {
		s += "Fizz"
	}
	if (i % 5) == 0 {
		s += "Buzz"
	}
	if s == "" {
		s = strconv.Itoa(i)
	}
	return s
}
