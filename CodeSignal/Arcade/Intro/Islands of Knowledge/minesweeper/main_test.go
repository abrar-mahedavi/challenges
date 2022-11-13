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
func Equals(a, b [][]int) bool {
	// fmt.Printf("%T %T", a, b)
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		// fmt.Printf("v => %v b => %v\n", v, b[i])
		//fmt.Printf("v(Type) => %T b(Type) => %T\n", v, b[i])
		// fmt.Printf("v != b[i] => %v \n", v != b[i])
		if !Equal(v, b[i]) {
			return false
		}
	}
	return true
}

func TestMinesweeperSingel(t *testing.T) {
	ans := minesweeper([][]bool{{true, false, false},
		{false, true, false},
		{false, false, false}})
	if !Equals(ans, [][]int{{1, 2, 1},
		{2, 1, 1},
		{1, 1, 1}}) {
		t.Errorf("minesweeper([][]bool{{true, false, false}, {false, true, false}, {false, false, false}} = %v; want [[1 2 1] [2 1 1] [1 1 1]]", ans)
	}
}

func TestMinesweeperTableDriven(t *testing.T) {
	var tests = []struct {
		a    [][]bool
		want [][]int
	}{
		{[][]bool{{true, false, false},
			{false, true, false},
			{false, false, false}}, [][]int{{1, 2, 1},
			{2, 1, 1},
			{1, 1, 1}}},
		{[][]bool{{false, false, false},
			{false, false, false}}, [][]int{{0, 0, 0},
			{0, 0, 0}}},
		{[][]bool{{true, false, false, true},
			{false, false, true, false},
			{true, true, false, true}}, [][]int{{0, 2, 2, 1},
			{3, 4, 3, 3},
			{1, 2, 3, 1}}},
		{[][]bool{{true, true, true},
			{true, true, true},
			{true, true, true}}, [][]int{{3, 5, 3},
			{5, 8, 5},
			{3, 5, 3}}},
		{[][]bool{{false, true, true, false},
			{true, true, false, true},
			{false, false, true, false}}, [][]int{{3, 3, 3, 2},
			{2, 4, 5, 2},
			{2, 3, 2, 2}}},
		{[][]bool{{true, false},
			{true, false},
			{false, true},
			{false, false},
			{false, false}}, [][]int{{1, 2},
			{2, 3},
			{2, 1},
			{1, 1},
			{0, 0}}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := minesweeper(tt.a)
			if !Equals(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
