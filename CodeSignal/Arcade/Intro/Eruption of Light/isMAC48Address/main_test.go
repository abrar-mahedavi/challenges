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

func TestIsMAC48AddressSingel(t *testing.T) {
	ans := isMAC48Address("00-1B-63-84-45-E6")
	if ans != true {
		t.Errorf("isMAC48Address(\"00-1B-63-84-45-E6\") = %v; want true", ans)
	}
}

func TestIsMAC48AddressTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want bool
	}{
		{"00-1B-63-84-45-E6", true},
		{"Z1-1B-63-84-45-E6", false},
		{"not a MAC-48 address", false},
		{"FF-FF-FF-FF-FF-FF", true},
		{"00-00-00-00-00-00", true},
		{"G0-00-00-00-00-00", false},
		{"12-34-56-78-9A-BC", true},
		{"02-03-04-05-06-07-", false},
		{"FF-FF-AB-CD-EA-BC", true},
		{"12-34-56-78-9A-BG", false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := isMAC48Address(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
