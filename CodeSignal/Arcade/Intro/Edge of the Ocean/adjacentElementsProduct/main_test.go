package main

import (
	"fmt"
	"testing"
)

func TestCenturyFromYearSingel(t *testing.T) {
	ans := centuryFromYear(2001)
	if ans != 21 {
		t.Errorf("centuryFromYear(2001) = %d; want 21", ans)
	}
}

func TestCenturyFromYearTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{1, 1},
		{1700, 17},
		{2001, 21},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := centuryFromYear(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
