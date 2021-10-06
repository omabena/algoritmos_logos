package main

import (
	"fmt"
	"testing"
)

func TestProduct(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{10, 10, 100},
		{5, 10, 50},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d * %d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := product(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}

func TestPower(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{4, 2, 16},
		{25, 10, 95367431640625},
		{4, 9, 262144},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d^%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := power(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}

func TestMod(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{10, 2, 0},
		{20, 7, 6},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%dmod%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := mod(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}

func TestDivInteger(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{10, 2, 5},
		{2, 10, 0},
		{20, 2, 10},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d divint %d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := div(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}

func TestSqrtGuess(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{25, 5},
		{10, -1},
		{2500, 50},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("sqrt %d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := sqrt(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}

func TestSqrtBigNumberGuess(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{25, 5},
		{10, -1},
		{2500, 50},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("sqrt %d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := sqrtBigNumber(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}

func TestCopyArray(t *testing.T) {
	var tests = []struct {
		a    []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("copyArray %d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := copyArray(tt.a)
			if !Equal(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
