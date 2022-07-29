package main

import "sort"

func maxProductDifference1(nums []int) int {
	max := nums[0] * nums[1]
	min := nums[0] * nums[1]

    for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			t := nums[i] * nums[j]

			if t > max {
				max = t
			}

			if t < min {
				min = t
			}
		}
	}

	return max - min
}

func maxProductDifference2(nums []int) int {
	sort.Ints(nums)

	return nums[len(nums) - 1] *  nums[len(nums) - 2] - nums[0] *  nums[1]
}
