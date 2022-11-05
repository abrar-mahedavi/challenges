package main

import (
	"fmt"
	"testing"
)

func Equal(a, b []string) bool {
	// fmt.Printf("%T %T", a, b)
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		// fmt.Printf("v => %v b => %v\n", v, b[i])
		// fmt.Printf("v(Type) => %T b(Type) => %T\n", v, b[i])
		// fmt.Printf("v == b[i] => %v \n", v == b[i])
		if v == b[i] {
			return false
		}
	}
	return true
}

func TestAllLongestStringsSingel(t *testing.T) {
	ans := allLongestStrings([]string{"aba",
		"aa",
		"ad",
		"vcd",
		"aba"})
	if Equal(ans, []string{"aba", "vcd", "aba"}) {
		t.Errorf("allLongestStrings([]string{\"aba\", \"aa\",\"ad\", \"vcd\", \"aba\"})= %v; want []string{\"aba\", \"vcd\", \"aba\"}", ans)
	}
}

func TestAllLongestStringsTableDriven(t *testing.T) {
	var tests = []struct {
		a    []string
		want []string
	}{
		{[]string{"aba",
			"aa",
			"ad",
			"vcd",
			"aba"}, []string{"aba",
			"vcd",
			"aba"}},
		{[]string{"aa"}, []string{"aa"}},
		{[]string{"abc",
			"eeee",
			"abcd",
			"dcd"}, []string{"eeee",
			"abcd"}},
		{[]string{"zzzzzz",
			"abcdef",
			"aaaaaa"}, []string{"zzzzzz",
			"abcdef",
			"aaaaaa"}},
		{[]string{"enyky",
			"benyky",
			"yely",
			"varennyky"}, []string{"varennyky"}},
		{[]string{"abacaba",
			"abacab",
			"abac",
			"xxxxxx"}, []string{"abacaba"}},
		{[]string{"young",
			"yooooooung",
			"hot",
			"or",
			"not",
			"come",
			"on",
			"fire",
			"water",
			"watermelon"}, []string{"yooooooung",
			"watermelon"}},
		{[]string{"onsfnib",
			"aokbcwthc",
			"jrfcw"}, []string{"aokbcwthc"}},
		{[]string{"lbgwyqkry"}, []string{"lbgwyqkry"}},
		{[]string{"i"}, []string{"i"}},
		{[]string{" "}, []string{" "}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := allLongestStrings(tt.a)
			if Equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
