package main

import "fmt"

//Input: numbers = [2,7,11,15], target = 9
//Output: [1,2]

func main() {
	inputNums := []int{2, 7, 11, 15}
	target := 9
	retIndex := sum(inputNums, target)
	fmt.Println(retIndex)
}

func sum(sortedNums []int, target int) []int {
	var idx = 0
	var lastIdx = len(sortedNums) - 1
	for {
		if idx == lastIdx {
			break
		} else if sortedNums[idx]+sortedNums[lastIdx] == target {
			break
		} else if sortedNums[idx]+sortedNums[lastIdx] > target {
			lastIdx -= 1
		} else if sortedNums[idx]+sortedNums[lastIdx] < target {
			idx += 1
		}
	}
	return []int{idx + 1, lastIdx + 1}
}
