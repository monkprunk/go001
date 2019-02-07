package money

import (
	"fmt"
	"testing"
)

func TestReturnMoney(t *testing.T) {
	testCases := []struct {
		input    int
		expected []Money
	}{
		{input: 555, expected: []Money{{value: 1000, amount: 0}, {value: 500, amount: 1},
			{value: 100, amount: 0}, {value: 50, amount: 1}, {value: 10, amount: 0},
			{value: 5, amount: 1}, {value: 2, amount: 0}, {value: 1, amount: 0}}},
		{input: 9999, expected: []Money{{value: 1000, amount: 9}, {value: 500, amount: 1},
			{value: 100, amount: 4}, {value: 50, amount: 1}, {value: 10, amount: 4},
			{value: 5, amount: 1}, {value: 2, amount: 2}, {value: 1, amount: 0}}},
		{input: 0, expected: []Money{{value: 1000, amount: 0}, {value: 500, amount: 0},
			{value: 100, amount: 0}, {value: 50, amount: 0}, {value: 10, amount: 0},
			{value: 5, amount: 0}, {value: 2, amount: 0}, {value: 1, amount: 0}}},
		{input: 1111, expected: []Money{{value: 1000, amount: 1}, {value: 500, amount: 0},
			{value: 100, amount: 1}, {value: 50, amount: 0}, {value: 10, amount: 1},
			{value: 5, amount: 0}, {value: 2, amount: 0}, {value: 1, amount: 1}}},
	}

	for _, v := range testCases {
		name := fmt.Sprintf("TestReturnMoney Input %v get %v", v.input, v.expected)
		t.Run(name, func(t *testing.T) {
			s := ReturnMoney(v.input)
			if (s[0].amount != v.expected[0].amount) || (s[1].amount != v.expected[1].amount) ||
				(s[2].amount != v.expected[2].amount) || (s[3].amount != v.expected[3].amount) ||
				(s[4].amount != v.expected[4].amount) || (s[5].amount != v.expected[5].amount) ||
				(s[6].amount != v.expected[6].amount) || (s[7].amount != v.expected[7].amount) {
				t.Errorf("should be %v but got %v", v.expected, s)
			}
		})
	}
}
