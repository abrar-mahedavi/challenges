package main

import (
	"fmt"
	"testing"
)

func Equal(a, b []int) bool {
	// fmt.Printf("%T %T", a, b)
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		// fmt.Printf("v => %v b => %v\n", v, b[i])
		//fmt.Printf("v(Type) => %T b(Type) => %T\n", v, b[i])
		// fmt.Printf("v != b[i] => %v \n", v != b[i])
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSumUpNumbers(t *testing.T) {
	ans := sumUpNumbers("13:58")
	if ans != true {
		t.Errorf("sumUpNumbers(\"2 apples and 12 oranges\") = %v; want \"14\"", ans)
	}
}

func TestSumUpNumbersDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want int
	}{
		{"2 apples, 12 oranges", 14},
		{"123450", 123450},
		{"Your payment method is invalid", 0},
		{"no digits at all", 0},
		{"there are some (12) digits 5566 in this 770 string 239", 6587},
		{"42+781", 823},
    {"abc abc 4 abc 0 abc", 4},
    {"abcdefghijklmnopqrstuvwxyz1AbCdEfGhIjKlMnOpqrstuvwxyz23,74 -", 98},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := sumUpNumbers(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
