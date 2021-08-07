package _97

func findJudge(n int, trust [][]int) int {
	var all []int
	var judgeSlice []int
	for _, x := range trust {
		if containsInt(all, x[0]) == -1 {
			all = append(all, x[0])
		}
		if containsInt(all, x[1]) == -1 {
			all = append(all, x[1])
		}
		var visited []int
		v := findTerminatedNode(trust, x[1], visited)
		if containsInt(judgeSlice, v) == -1 && v != -1 {
			judgeSlice = append(judgeSlice, v)
		}

	}
	if len(all) == n && len(judgeSlice) == 1 {
		return judgeSlice[0]
	}
	return -1
}

func findTerminatedNode(arr [][]int, val int, visited []int) int {
	if len(arr) == len(visited) {
		return -1
	}
	visited = append(visited, val)
	for _, ns := range arr {
		if ns[0] == val {

			return findTerminatedNode(arr, ns[1], visited)
		}
	}
	return val
}

func containsInt(arr []int, val int) int {
	for idx, v := range arr {
		if v == val {
			return idx
		}
	}
	return -1
}
