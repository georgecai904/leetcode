package main

import "fmt"

//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice.

func filter(nums []int, callback func(int) bool) []int {
	var results []int
	for _, val := range nums {
		if callback(val) {
			results = append(results, val)
		} else {
			results = append(results, -1)
		}
	}
	return results
}

func twoSum(nums []int, target int) []int {

	nums = filter(nums, func(val int) bool {
		return val <= target
	})
	var result []int
	for i, val := range nums {
		temp := target - val
		for j, val2 := range nums[i+1:] {
			if val2 == temp {
				result = append(result, i)
				result = append(result, i+j)
				fmt.Println(val2, val, i, j)
				break
			}
		}
	}
	return result
}
func main() {
	nums := []int{0, 3, 2, 0}
	target := 0

	fmt.Println(twoSum(nums, target))
}
