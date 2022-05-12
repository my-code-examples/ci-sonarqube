package main

import "testing"

type Test struct {
	i1       int
	i2       int
	expected int
}

type Tests []Test

func TestSum(t *testing.T) {
	tests := Tests{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{-1, -1, -2},
	}
	for _, test := range tests {
		if sum := sum(test.i1, test.i2); sum != test.expected {
			t.Errorf("sum(%d, %d) returned %d, expected %d",
				test.i1, test.i2, sum, test.expected)
		}
	}
}

// func TestSub(t *testing.T) {
// 	tests := Tests{
// 		{1, 2, -1},
// 		{0, 0, 0},
// 		{-1, 1, -2},
// 		{-1, -1, 0},
// 	}
// 	for _, test := range tests {
// 		if sub := sub(test.i1, test.i2); sub != test.expected {
// 			t.Errorf("sub(%d, %d) returned %d, expected %d",
// 				test.i1, test.i2, sub, test.expected)
// 		}
// 	}
// }

// func TestTimes(t *testing.T) {
// 	tests := Tests{
// 		{1, 2, 2},
// 		{0, 0, 0},
// 		{-1, 1, -1},
// 		{-1, -1, 1},
// 	}
// 	for _, test := range tests {
// 		if times := times(test.i1, test.i2); times != test.expected {
// 			t.Errorf("times(%d, %d) returned %d, expected %d",
// 				test.i1, test.i2, times, test.expected)
// 		}
// 	}
// }

// func TestDiv(t *testing.T) {
// 	tests := Tests{
// 		{1, 2, 0},
// 		{0, 0, 0},
// 		{-1, 1, -1},
// 		{-1, -1, 1},
// 	}
// 	for _, test := range tests {
// 		if div := div(test.i1, test.i2); div != test.expected {
// 			t.Errorf("div(%d, %d) returned %d, expected %d",
// 				test.i1, test.i2, div, test.expected)
// 		}
// 	}
// }
