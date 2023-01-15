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

func TestSpiralNumbersSingle(t *testing.T) {
	ans := spiralNumbers(3)
	if !Equals(ans, [][]int{{1,2,3},
                            {8,9,4},
                            {7,6,5}}) {
		t.Errorf("spiralNumbers(3) = %v; want [[1 2 3] [8 9 4] [7 6 5]]", ans)
	}
}

func TestSpiralNumbersTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		want [][]int
	}{
		{3, [][]int{{1,2,3},
                  {8,9,4},
                  {7,6,5}}},
    {5, [][]int{{1,2,3,4,5},
                  {16,17,18,19,6},
                  {15,24,25,20,7},
                  {14,23,22,21,8},
                  {13,12,11,10,9}}},
    {6, [][]int{{1,2,3,4,5,6},
                  {20,21,22,23,24,7},
                  {19,32,33,34,25,8},
                  {18,31,36,35,26,9},
                  {17,30,29,28,27,10},
                  {16,15,14,13,12,11}}},
    {7, [][]int{{1,2,3,4,5,6,7},
                  {24,25,26,27,28,29,8},
                  {23,40,41,42,43,30,9},
                  {22,39,48,49,44,31,10},
                  {21,38,47,46,45,32,11},
                  {20,37,36,35,34,33,12},
                  {19,18,17,16,15,14,13}}},
    {10, [][]int{{1,2,3,4,5,6,7,8,9,10},
                   {36,37,38,39,40,41,42,43,44,11},
                   {35,64,65,66,67,68,69,70,45,12},
                   {34,63,84,85,86,87,88,71,46,13},
                   {33,62,83,96,97,98,89,72,47,14},
                   {32,61,82,95,100,99,90,73,48,15},
                   {31,60,81,94,93,92,91,74,49,16},
                   {30,59,80,79,78,77,76,75,50,17},
                   {29,58,57,56,55,54,53,52,51,18},
                   {28,27,26,25,24,23,22,21,20,19}}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := spiralNumbers(tt.a)
			if !Equals(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
