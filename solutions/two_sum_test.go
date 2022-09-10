package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9
	rs := twoSum(input, target)
	fmt.Println(rs)
}
