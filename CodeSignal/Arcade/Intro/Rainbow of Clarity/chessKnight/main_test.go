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

func TestChessKnightSingel(t *testing.T) {
	ans := chessKnight("a1")
	if ans != 2 {
		t.Errorf("chessKnight(\"a1\") = %v; want 2", ans)
	}
}

func TestChessKnightTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want int
	}{
		{"a1", 2},
		{"c2", 6},
		{"d4", 8},
		{"g6", 6},
		{"a3", 4},
		{"b7", 4},
		{"h8", 2},
		{"h6", 4},
		{"g8", 3},
		{"a5", 4},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := chessKnight(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
