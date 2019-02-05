package tdd

import (
	"fmt"
	"testing"
)

func TestSay(t *testing.T) {
	testCases := []struct {
		num      int
		expected string
	}{
		{num: 1, expected: "1"},
		{num: 2, expected: "2"},
		{num: 3, expected: "Fizz"},
		{num: 4, expected: "4"},
		{num: 5, expected: "Buzz"},
		{num: 6, expected: "Fizz"},
		{num: 7, expected: "7"},
		{num: 8, expected: "8"},
		{num: 9, expected: "Fizz"},
		{num: 10, expected: "Buzz"},
		{num: 12, expected: "Fizz"},
		{num: 13, expected: "13"},
		{num: 15, expected: "FizzBuzz"},
		{num: 30, expected: "FizzBuzz"},
	}

	for _, v := range testCases {
		name := fmt.Sprintf("TestSayInput %v get %v", v.num, v.expected)
		t.Run(name, func(t *testing.T) {
			s := Say(v.num)
			if s != v.expected {
				t.Errorf("should be %v but got %v", v.expected, s)
			}
		})
	}
}

func TestSayInput1Get1(t *testing.T) {
	s := Say(1)
	if s != "1" {
		t.Errorf("should be %v but got %v", "1", s)
	}
}

func BenchmarkSay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Say(15)
	}
}
