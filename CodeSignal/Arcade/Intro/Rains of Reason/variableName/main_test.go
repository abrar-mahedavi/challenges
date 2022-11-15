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

func TestVariableNameSingel(t *testing.T) {
	ans := variableName("var_1__Int")
	if ans != true {
		t.Errorf("variableName(\"var_1__Int\") = %v; want true", ans)
	}
}

func TestVariableNameTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want bool
	}{
		{"var_1__Int", true},
		{"qq-q", false},
		{"variable", true},
		{"2w2", false},
		{"va[riable0", false},
		{"variable0", true},
		{"a", true},
		{"_Aas_23", true},
		{"a a_2", false},
		{"0ss", false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := variableName(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
