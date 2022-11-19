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

func TestArrayMaxConsecutiveSumSingel(t *testing.T) {
	ans := arrayMaxConsecutiveSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
	if ans != 27 {
		t.Errorf("arrayMaxConsecutiveSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3) = %v; want 27", ans)
	}
}

func TestArrayMaxConsecutiveSumTableDriven(t *testing.T) {
	var tests = []struct {
		a    []int
		b    int
		want int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 27},
		{[]int{2, 4, 10, 1}, 2, 14},
		{[]int{1, 3, 2, 4}, 3, 9},
		{[]int{3, 2, 1, 1}, 1, 3},
		{[]int{1, 3, 4, 2, 4, 2, 4}, 4, 13},
		// {[]int{768, 77, 755, 960, 747, 25, 107, 520, 995, 404, 43, 714, 632, 642, 493, 352, 450, 625, 262, 217, 254, 55, 661, 822, 562, 187, 603, 216, 275, 76, 75, 417, 350, 942, 1000, 232, 887, 173, 858, 116, 75, 170, 529, 26, 62, 378, 667, 444, 240, 325, 444, 391, 698, 282, 870, 611, 974, 388, 586, 616, 845, 591, 525, 976, 938, 673, 413, 862, 396, 856, 764, 415, 309, 27, 583, 630, 741, 988, 456, 807, 242, 624, 149, 524, 962, 960, 900, 199, 645, 36, 343, 943, 232, 781, 445, 670, 177, 889, 57, 519}, 250, 135845},
		{[]int{963, 741, 22, 851, 399, 382, 190, 247, 494, 452, 891, 532, 795, 920, 465, 831, 344, 391, 582, 897, 763, 951, 735, 806, 320, 702, 200, 59, 870, 345, 695, 321, 817, 83, 416, 36, 914, 335, 117, 985, 690, 303, 875, 556, 292, 309, 496, 791, 509, 360, 235, 784, 607, 341}, 23, 14232},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := arrayMaxConsecutiveSum(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
