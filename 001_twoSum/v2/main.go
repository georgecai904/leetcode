package main

import "fmt"

func twoSum(nums []int, target int) []int {
	d := make(map[int]int)
	var result []int
	for i, val := range nums {
		if _, exists := d[val]; exists {
			result = []int{d[val], i}
		} else {
			d[target-val] = i
		}
	}
	return result
}

func main() {
	nums := []int{3, 2, 4}
	target := 6

	fmt.Println(twoSum(nums, target))
}
