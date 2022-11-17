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

func TestDepositProfitSingel(t *testing.T) {
	ans := depositProfit(100, 20, 170)
	if ans != 3 {
		t.Errorf("chessBoardCellColor(\"A1\", \"C3\") = %v; want true", ans)
	}
}

func TestDepositProfitTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		c    int
		want int
	}{
		{100, 20, 170, 3},
		{100, 1, 101, 1},
		{1, 100, 64, 6},
		{20, 20, 50, 6},
		{50, 25, 100, 4},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v %v", tt.a, tt.b, tt.c)
		t.Run(testname, func(t *testing.T) {
			ans := depositProfit(tt.a, tt.b, tt.c)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
