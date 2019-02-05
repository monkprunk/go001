package tdd

import "strconv"

func Say(i int) string {

	s := []string{"FizzBuzz", "x", "x", "Fizz", "x", "Buzz", "Fizz", "x",
		"x", "Fizz", "Buzz", "x", "Fizz", "x", "x"}
	var result string
	result = s[i%15]
	if result == "x" {
		return strconv.Itoa(i)
	}
	return result
}
