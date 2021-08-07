package _97

import "testing"

func Test_findJudge(t *testing.T) {
	type args struct {
		n     int
		trust [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 2, trust: [][]int{{1, 2}}}, want: 2},
		{name: "2", args: args{n: 3, trust: [][]int{{1, 3}, {2, 3}}}, want: 3},
		{name: "3", args: args{n: 3, trust: [][]int{{1, 3}, {2, 3}, {3, 1}}}, want: -1},
		{name: "4", args: args{n: 3, trust: [][]int{{1, 2}, {2, 3}}}, want: 3},
		{name: "5", args: args{n: 4, trust: [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findJudge(tt.args.n, tt.args.trust); got != tt.want {
				t.Errorf("findJudge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findTerminatedNode(t *testing.T) {
	type args struct {
		arr [][]int
		val int
		visited []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{arr: [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}, val: 1, visited: []int{}}, want: 3},
		{name: "2", args: args{arr: [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}, val: 2, visited: []int{}}, want: 3},
		{name: "3", args: args{arr: [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}, val: 3, visited: []int{}}, want: 3},
		{name: "4", args: args{arr: [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}, val: 4, visited: []int{}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTerminatedNode(tt.args.arr, tt.args.val, tt.args.visited); got != tt.want {
				t.Errorf("findTerminatedNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
