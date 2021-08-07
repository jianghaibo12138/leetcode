package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var maxLen int
	var minIdx = 0
	var curLen = 0
	var visited []byte
	var bytes = []byte(s)
	for _, b := range bytes {
		curLen += 1
		for {
			index := containsByte(&visited, b)
			if index == -1 {
				break
			}
			index = containsByte(&visited, s[minIdx])
			visited = append(visited[:index], visited[index+1:]...)
			minIdx += 1
			curLen -= 1
		}
		if curLen > maxLen {
			maxLen = curLen
		}
		visited = append(visited, b)
	}
	return maxLen
}

func containsByte(bs *[]byte, n byte) int {
	var index = -1
	for idx, b := range *bs {
		if b == n {
			return idx
		}
	}
	return index
}

func main() {
	s := "abcabcbb"
	maxLen := lengthOfLongestSubstring(s)
	fmt.Println(maxLen)
}
