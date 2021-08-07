package main

import "fmt"

func main() {
	inputNums := []int{-1, -2, -3, -4, -5}
	target := -8
	retIndex := twoSum(inputNums, target)
	fmt.Println(retIndex)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	num := 0
	for idx, val := range nums {
		num = target - val
		index, ok := m[num]
		if ok {
			return []int{index, idx}
		}
		m[val] = idx
	}
	return []int{-1, -1}
}
