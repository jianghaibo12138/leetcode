package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var targetNums = nums1
	for _, n := range nums1 {
		for _, m := range nums2 {
			if m <= n {
				idx := containsInt(targetNums, n)
				targetNums = insert(targetNums, idx, m)
			}
		}
	}
	fmt.Println(targetNums)
	return 0
}

func containsInt(nums []int, val int) int {
	for idx, n := range nums {
		if n == val {
			return idx
		}
	}
	return -1
}

func insert(nums []int, index, val int) []int {
	ns := nums[:index]
	ns = append(ns, val)
	ns = append(ns, nums[index+1:]...)
	return ns
}
