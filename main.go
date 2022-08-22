package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findMaxConsecutiveOnes(nums []int) int {
    var result int
    var max int
    
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            max += nums[i]
        } else if nums[i] == 0 {
			max = 0
        }

		if max > result {
			result = max
		}
    }
    
    return result
}

func findNumbers(nums []int) int {
    var result int
    
    for _, n := range nums {
        if len(strconv.Itoa(n))%2 == 0 {
            result++
        }
    }
    
    return result
}

func AbcInt(n int) int {
    if n < 0 {
        n *= -1
        return n
    }

    return n
}


func sortedSquares(nums []int) []int {
    var result []int

    for i := 0; i < len(nums); i++ {
        n := AbcInt(nums[i])

        result = append(result, n * n)
    }

    sort.Ints(result)

    return result
}

func main() {
	fmt.Println(sortedSquares([]int{-4,-1,0,3,10}))
}
