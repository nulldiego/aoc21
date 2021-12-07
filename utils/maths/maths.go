package maths

import "sort"

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Median(nums []int) int {
	sort.Ints(nums)

	half := len(nums) / 2

	if len(nums)%2 != 0 {
		return nums[half]
	}

	return (nums[half-1] + nums[half]) / 2
}

func Mean(nums []int) float64 {
	total := 0

	for _, n := range nums {
		total += n
	}

	return float64(total) / float64(len(nums))
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
