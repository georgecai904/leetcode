//There are two sorted arrays nums1 and nums2 of size m and n respectively.
//Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//var totalLength = len(nums1) + len(nums2)
	//for
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{4, 5}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
