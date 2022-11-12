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

func TestIsIPv4AddressSingel(t *testing.T) {
	ans := isIPv4Address("10.15.15.10")
	if ans != true {
		t.Errorf("isIPv4Address(\"10.15.15.10\") = %v; want true", ans)
	}
}

func TestIsIPv4AddressTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want bool
	}{
		{"172.16.254.1", true},
		{"172.316.254.1", false},
		{".254.255.0", false},
		{"1.1.1.1a", false},
		{"1", false},
		{"0.254.255.0", true},
		{"1.23.256.255.", false},
		{"1.23.256..", false},
		{"0..1.0", false},
		{"64.233.161.00", false},
		{"64.00.161.131", false},
		{"01.233.161.131", false},
		{"35..36.9.9.0", false},
		{"1.1.1.1.1", false},
		{"1.256.1.1", false},
		{"a0.1.1.1", false},
		{"0.1.1.256", false},
		{"129380129831213981.255.255.255", false},
		{"255.255.255.255abcdekjhf", false},
		{"7283728", false},
		{"0..1.0.0", false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := isIPv4Address(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
