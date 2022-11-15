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

func TestChessBoardCellColorSingel(t *testing.T) {
	ans := chessBoardCellColor("A1", "C3")
	if ans != true {
		t.Errorf("chessBoardCellColor(\"A1\", \"C3\") = %v; want true", ans)
	}
}

func TestChessBoardCellColorTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		b    string
		want bool
	}{
		{"C3", "A1", true},
		{"A1", "H3", false},
		{"A1", "A2", false},
		{"A1", "A2", false},
		{"B3", "H8", false},
		{"C3", "B5", false},
		{"G5", "E7", true},
		{"C8", "H8", false},
		{"D2", "D2", true},
		{"A2", "A5", false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := chessBoardCellColor(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
