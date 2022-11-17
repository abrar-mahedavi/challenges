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

func TestStringsRearrangementSingel(t *testing.T) {
	ans := stringsRearrangement([]string{"aba",
		"bbb",
		"bab"})
	if ans != false {
		t.Errorf("chessBoardCellColor(\"A1\", \"C3\") = %v; want true", ans)
	}
}

func TestStringsRearrangementTableDriven(t *testing.T) {
	var tests = []struct {
		a    []string
		want bool
	}{
		{[]string{"aba",
			"bbb",
			"bab"}, false},
		{[]string{"ab",
			"bb",
			"aa"}, true},
		{[]string{"q",
			"q"}, false},
		{[]string{"zzzzab",
			"zzzzbb",
			"zzzzaa"}, true},
		{[]string{"ab",
			"ad",
			"ef",
			"eg"}, false},
		{[]string{"abc",
			"bef",
			"bcc",
			"bec",
			"bbc",
			"bdc"}, true},
		{[]string{"abc",
			"abx",
			"axx",
			"abc"}, false},
		{[]string{"abc",
			"abx",
			"axx",
			"abx",
			"abc"}, false},
		{[]string{"f",
			"g",
			"a",
			"h"}, true},
		{[]string{"ff",
			"gf",
			"af",
			"ar",
			"hf"}, true},
		{[]string{"a",
			"b",
			"c"}, true},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := stringsRearrangement(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
