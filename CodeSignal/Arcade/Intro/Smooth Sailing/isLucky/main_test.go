package main

import (
	"fmt"
	"testing"
)

func TestIsLuckySingel(t *testing.T) {
	ans := isLucky(1230)
	if ans != true {
		t.Errorf("isLucky(1230) = %v; want true", ans)
	}
}

func TestIsLuckyTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		want bool
	}{
		{1230, true},
		{239017, false},
		{134008, true},
		{10, false},
		{11, true},
		{1010, true},
		{261534, false},
		{100000, false},
		{999999, true},
		{123321, true},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := isLucky(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
