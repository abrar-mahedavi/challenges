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

func TestFirstDigitSingel(t *testing.T) {
	ans := firstDigit("var_1__Int")
	if ans != "1" {
		t.Errorf("firstDigit(\"var_1__Int\") = %v; want \"1\" ", ans)
	}
}

func TestFirstDigitTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want string
	}{
		{"dasd_1_adsds", "1"},
		{"q2q-q", "2"},
		{"0ss", "0"},
		{"_Aas_23", "2"},
		{"a a_933", "9"},
		{"ok0", "0"},
		{"abu", " "},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := firstDigit(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
