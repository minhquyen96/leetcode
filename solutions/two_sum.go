package leetcode

func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	for i, e := range nums {
		for j := i + 1; j < len(nums); j++ {
			if e+nums[j] == target {
				res = append(res, i, j)
				return res
			}
		}
	}
	return res
}
