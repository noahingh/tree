package tree

import (
	"reflect"
	"testing"
)

func Test_topologicalSort(t *testing.T) {
	type args struct {
		adj [][]bool
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "circuit",
			args: args{
				adj: [][]bool{
					{false, true},
					{true, false},
				},
			},
			want:    []int{},
			wantErr: true,
		},
		{
			name: "step",
			args: args{
				adj: [][]bool{
					{false, true, false, false},
					{false, false, true, false},
					{false, false, false, true},
					{false, false, false, false},
				},
			},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "file system",
			args: args{
				adj: [][]bool{
					{false, true, false, false, true, false}, // 0 -> 1,4
					{false, false, true, true, false, false}, // 1 -> 2,3
					{false, false, false, false, false, false},
					{false, false, false, false, false, false},
					{false, false, false, false, false, true}, // 4 -> 5
					{false, false, false, false, false, false},
				},
			},
			// the result of dfs is [2, 3, 1, 5, 4, 0]
			want: []int{0, 4, 5, 1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := topologicalSort(tt.args.adj)
			if (err != nil) != tt.wantErr {
				t.Errorf("topologicalSort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topologicalSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasCircuit(t *testing.T) {
	type args struct {
		sorted []int
		adj    [][]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "circuit exist",
			args: args{
				sorted: []int{0, 1},
				adj: [][]bool{
					{false, true},
					{true, false},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCircuit(tt.args.sorted, tt.args.adj); got != tt.want {
				t.Errorf("hasCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet []int
	}{
		{
			name: "increasing slice",
			args: args{
				ints: []int{0, 1, 2, 3, 4},
			},
			wantRet: []int{4, 3, 2, 1, 0},
		},
		{
			name: "decreasing slice",
			args: args{
				ints: []int{4, 3, 2, 1, 0},
			},
			wantRet: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := reverse(tt.args.ints); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("reverse() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
