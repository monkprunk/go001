package roman

import (
	"fmt"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	testCases := []struct {
		num      int
		expected string
	}{
		{num: 1, expected: "I"},
		{num: 2, expected: "II"},
		{num: 3, expected: "III"},
		{num: 4, expected: "IV"},
		{num: 5, expected: "V"},
		{num: 6, expected: "VI"},
		{num: 7, expected: "VII"},
		{num: 8, expected: "VIII"},
		{num: 9, expected: "IX"},
		{num: 10, expected: "X"},
		{num: 12, expected: "XII"},
		{num: 14, expected: "XIV"},
		{num: 15, expected: "XV"},
		{num: 40, expected: "XL"},
		{num: 44, expected: "XLIV"},
		{num: 50, expected: "L"},
		{num: 60, expected: "LX"},
		{num: 70, expected: "LXX"},
		{num: 80, expected: "LXXX"},
		{num: 90, expected: "XC"},
		{num: 99, expected: "XCIX"},
		{num: 100, expected: "C"},
		{num: 999, expected: "CMXCIX"},
	}

	for _, v := range testCases {
		name := fmt.Sprintf("TestIntToRomanInput %v get %v", v.num, v.expected)
		t.Run(name, func(t *testing.T) {
			s := IntToRoman(v.num)
			if s != v.expected {
				t.Errorf("should be %v but got %v", v.expected, s)
			}
		})
	}
}
