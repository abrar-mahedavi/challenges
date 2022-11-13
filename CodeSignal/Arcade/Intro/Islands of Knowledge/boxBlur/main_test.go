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

func TestBoxBlurSingel(t *testing.T) {
	ans := boxBlur([][]int{{1, 1, 1},
		{1, 7, 1},
		{1, 1, 1}})
	if !Equals(ans, [][]int{{1}}) {
		t.Errorf("boxBlur([][]int{{1, 1, 1},{1, 7, 1},{1, 1, 1}}) = %v; want [[1]]", ans)
	}
}

func TestBoxBlurTableDriven(t *testing.T) {
	var tests = []struct {
		a    [][]int
		want [][]int
	}{
		{[][]int{{1, 1, 1},
			{1, 7, 1},
			{1, 1, 1}}, [][]int{{1}}},
		{[][]int{{0, 18, 9},
			{27, 9, 0},
			{81, 63, 45}}, [][]int{{28}}},
		{[][]int{{36, 0, 18, 9},
			{27, 54, 9, 0},
			{81, 63, 72, 45}}, [][]int{{40, 30}}},
		{[][]int{{7, 4, 0, 1},
			{5, 6, 2, 2},
			{6, 10, 7, 8},
			{1, 4, 2, 0}}, [][]int{{5, 4}, {4, 4}}},
		{[][]int{{36, 0, 18, 9, 9, 45, 27},
			{27, 0, 54, 9, 0, 63, 90},
			{81, 63, 72, 45, 18, 27, 0},
			{0, 0, 9, 81, 27, 18, 45},
			{45, 45, 27, 27, 90, 81, 72},
			{45, 18, 9, 0, 9, 18, 45},
			{27, 81, 36, 63, 63, 72, 81}}, [][]int{{39, 30, 26, 25, 31},
			{34, 37, 35, 32, 32},
			{38, 41, 44, 46, 42},
			{22, 24, 31, 39, 45},
			{37, 34, 36, 47, 59}}},
		{[][]int{{36, 0, 18, 9, 9, 45, 27},
			{27, 0, 254, 9, 0, 63, 90},
			{81, 255, 72, 45, 18, 27, 0},
			{0, 0, 9, 81, 27, 18, 45},
			{45, 45, 227, 227, 90, 81, 72},
			{45, 18, 9, 255, 9, 18, 45},
			{27, 81, 36, 127, 255, 72, 81}}, [][]int{{82, 73, 48, 25, 31},
			{77, 80, 57, 32, 32},
			{81, 106, 88, 68, 42},
			{44, 96, 103, 89, 45},
			{59, 113, 137, 126, 80}}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := boxBlur(tt.a)
			if !Equals(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
