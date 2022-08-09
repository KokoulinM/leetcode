package main

import "sort"

func maxProduct1(nums []int) int {
	var sum int

    for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			tmp := (nums[i]-1)*(nums[j]-1)

			if tmp > sum {
				sum = tmp
			}
		}
	}

	return sum
}

func maxProduct2(nums []int) int {
	sort.Ints(nums)

	return (nums[len(nums) - 2]-1)*(nums[len(nums) - 1]-1)
}