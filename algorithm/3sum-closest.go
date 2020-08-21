package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -2, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	sort.Ints(nums)
	sum := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums); i++ {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			current := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(sum-current)) > math.Abs(float64(target-current)) {
				sum = current
			}
			if current < target {
				l++
			} else if current == target {
				return target
			} else {
				r--
			}
		}
	}
	return sum
}
