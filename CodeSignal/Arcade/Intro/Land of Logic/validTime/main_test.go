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

func TestValidTime(t *testing.T) {
	ans := validTime("13:58")
	if ans != true {
		t.Errorf("validTime(\"13:58\") = %v; want \"true\"", ans)
	}
}

func TestValidTimeTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want bool
	}{
		{"13:58", true},
		{"25:51", false},
		{"02:76", false},
		{"24:00", false},
		{"02:61", false},
		{"27:00", false},
    {"19:66", false},
    {"12:31", true},
    {"25:73", false},
    {"09:56", true},
    {"03:29", true},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := validTime(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
