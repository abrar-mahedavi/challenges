package main

import (
	"fmt"
	"testing"
)

func Equal(a, b []string) bool {
	// fmt.Printf("%T %T", a, b)
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		// fmt.Printf("v => %v b => %v\n", v, b[i])
		// fmt.Printf("v(Type) => %T b(Type) => %T\n", v, b[i])
		// fmt.Printf("v == b[i] => %v \n", v == b[i])
		if v == b[i] {
			return false
		}
	}
	return true
}

func TestFileNamingSingle(t *testing.T) {
	ans := fileNaming([]string{"doc",
                              "doc",
                              "image",
                              "doc(1)",
                              "doc"})
	if Equal(ans, []string{"doc",
                          "doc(1)",
                          "image",
                          "doc(1)(1)",
                          "doc(2)"}) {
		t.Errorf("allLongestStrings([]string{\"doc\", \"doc\",\"image\",\"doc(1)\",\"doc\"})= %v; want []string{\"doc\",\"doc(1)\",\"image\",\"doc(1)(1)\",\"doc(2)\"}", ans)
	}
}

func TestFileNamingTableDriven(t *testing.T) {
	var tests = []struct {
		a    []string
		want []string
	}{
		{[]string{"doc",
               "doc",
               "image",
               "doc(1)",
               "doc"}, []string{"doc",
                                 "doc(1)",
                                 "image",
                                 "doc(1)(1)",
                                 "doc(2)"}},
		{[]string{"a(1)",
               "a(6)",
               "a",
               "a",
               "a",
               "a",
               "a",
               "a",
               "a",
               "a",
               "a",
               "a"}, []string{"a(1)",
                               "a(6)",
                               "a",
                               "a(2)",
                               "a(3)",
                               "a(4)",
                               "a(5)",
                               "a(7)",
                               "a(8)",
                               "a(9)",
                               "a(10)",
                               "a(11)"}},
		{[]string{"dd",
               "dd(1)",
               "dd(2)",
               "dd",
               "dd(1)",
               "dd(1)(2)",
               "dd(1)(1)",
               "dd",
               "dd(1)"}, []string{"dd",
                                   "dd(1)",
                                   "dd(2)",
                                   "dd(3)",
                                   "dd(1)(1)",
                                   "dd(1)(2)",
                                   "dd(1)(1)(1)",
                                   "dd(4)",
                                   "dd(1)(3)"}},
		{[]string{"a",
               "b",
               "cd",
               "b ",
               "a(3)"}, []string{"a",
                           "b",
                           "cd",
                           "b ",
                           "a(3)"}},
		{[]string{"name",
               "name",
               "name(1)",
               "name(3)",
               "name(2)",
               "name(2)"}, []string{"name",
                                     "name(1)",
                                     "name(1)(1)",
                                     "name(3)",
                                     "name(2)",
                                     "name(2)(1)"}},
		{[]string{}, []string{}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := allLongestStrings(tt.a)
			if Equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
