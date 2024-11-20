package main

import "testing"



func TestMysum(t *testing.T) {
	type test struct {
		data [] int
		answer int
	}
	
	tests := []test{
		{[]int{1, 2, 3},6},
		{[]int{1, 2, 3, 4, 5, 6, 7},28},
		{[]int{7, 8, 9, 10},34},	
		{[]int{-1,0,1},0},
	}
	
	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}
	}
}