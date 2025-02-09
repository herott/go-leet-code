package test

import (
	retry "algorithms/001-two-sum"
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := [][]int{
		{3, 2, 4},
		{3, 2, 4},
		{0, 8, 7, 3, 3, 4, 2},
		{0, 1},
	}
	targets := []int{
		6,
		5,
		11,
		1,
	}
	results := [][]int{
		{1, 2},
		{0, 1},
		{1, 3},
		{0, 1},
	}
	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")
	for i := 0; i < len(targets); i++ {
		fmt.Printf("nums = %v target = %v result = %v\n", tests[i], targets[i], retry.TwoSum(tests[i], targets[i]))
		if ret := retry.TwoSum(tests[i], targets[i]); ret[0] != results[i][0] && ret[1] != results[i][1] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
