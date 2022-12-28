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

func TestFindEmailDomainSingel(t *testing.T) {
	ans := findEmailDomain("prettyandsimple@example.com")
	if ans != "example.com" {
		t.Errorf("findEmailDomain(\"prettyandsimple@example.com\") = %v; want \"example.com\"", ans)
	}
}

func TestFindEmailDomainTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want string
	}{
		{"prettyandsimple@example.com", "example.com"},
		{"fully-qualified-domain@codesignal.com", "codesignal.com"},
		{"\" \"@space.com", "space.com"},
		{"someaddress@yandex.ru", "yandex.ru"},
		{"\" \"@xample.org", "xample.org"},
		{"\"much.more unusual\"@yahoo.com", "yahoo.com"},
		{"\"very.unusual.@.unusual.com\"@usual.com", "usual.com"},
		{"admin@mailserver2.ru", "mailserver2.ru"},
		{"example-indeed@strange-example.com", "strange-example.com"},
		{"very.common@gmail.com", "gmail.com"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := findEmailDomain(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
