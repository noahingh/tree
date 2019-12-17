package tree

import (
	"reflect"
	"testing"
)

func TestHasCircuit(t *testing.T) {
	type args struct {
		adj [CntItemLimit][CntItemLimit]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "circuit",
			args: args{
				adj: [CntItemLimit][CntItemLimit]bool{
					{false, true},
					{true, false},
				},
			},
			want: true,
		},
		{
			name: "step",
			args: args{
				adj: [CntItemLimit][CntItemLimit]bool{
					{false, true},
					{false, false, true},
					{false, false, false, true},
				},
			},
			want: false,
		},
		{
			name: "file system",
			args: args{
				adj: [CntItemLimit][CntItemLimit]bool{
					{false, true},
					{false, false, true, true}, // 1 -> 2,3
					{false, false, false, false},
					{false, false, false, false},
					{false, false, false, false, false, true}, // 4 -> 5
					{false, false, false, false, false, false},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasCircuit(tt.args.adj); got != tt.want {
				t.Errorf("HasCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_hasCircuit(t *testing.T) {
	type args struct {
		sorted []int
		adj    [CntItemLimit][CntItemLimit]bool
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
				adj: [CntItemLimit][CntItemLimit]bool{
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

