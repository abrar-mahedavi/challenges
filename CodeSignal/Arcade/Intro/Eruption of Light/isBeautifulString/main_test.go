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

func TestIsBeautifulStringSingel(t *testing.T) {
	ans := isBeautifulString("bbbaacdafe")
	if ans != true {
		t.Errorf("isBeautifulString(\"bbbaacdafe\") = %v; want true", ans)
	}
}

func TestIsBeautifulStringTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want bool
	}{
		{"bbbaacdafe", true},
		{"aabbb", false},
		{"bbc", false},
		{"bbbaa", false},
		{"abcdefghijklmnopqrstuvwxyzz", false},
		{"abcdefghijklmnopqrstuvwxyz", true},
		{"abcdefghijklmnopqrstuvwxyzqwertuiopasdfghjklxcvbnm", true},
		{"fyudhrygiuhdfeis", false},
		{"zaa", false},
		{"zyy", false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := isBeautifulString(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
