package tree

import (
	"reflect"
	"testing"
)

// NormalCase return a successful case.
// 0
// └── 1 
//     ├── 2 
// 	   └── 3
func NormalCase() (sorted []int, adj map[int][]int) {
	sorted = []int {0, 1, 3, 2}
	adj = map[int][]int{
		0: []int{1},
		1: []int{2, 3},
		2: []int{},
		3: []int{},
	}
	return sorted, adj
}

// CircuitCase return a circuit case.
// 0
// └── 1 
//     └── 0
func CircuitCase() (sorted []int, adj map[int][]int) {
	sorted = []int{0, 1}
	adj = map[int][]int{
		0: []int{1},
		1: []int{0},
	}
	return sorted, adj
}

// ComplexCase return a complex case.
// 0
// ├── 1
// │   ├── 2
// │   └── 3
// └── 4
//     └── 5
func ComplexCase() (sorted []int, adj map[int][]int) {
	sorted = []int{0, 4, 5, 1, 3, 2}
	adj = map[int][]int{
		0: []int{1, 4},
		1: []int{2, 3},
		2: []int{},
		3: []int{},
		4: []int{5},
		5: []int{},
	}

	return sorted, adj
}

func TestTopologicalSort(t *testing.T) {
	ns, na := NormalCase()
	_, ca := CircuitCase()
	ccs, cca := ComplexCase()

	type args struct {
		adj map[int][]int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "normal case",
			args: args{
				adj: na,
			},
			want: ns,
			wantErr: false,
		},
		{
			name: "circuit case",
			args: args{
				adj: ca,
			},
			want: []int{},		// should get a empty slice.
			wantErr: true,
		},
		{
			name: "complex case",
			args: args{
				adj: cca,
			},
			want: ccs,
			wantErr: false,
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

func TestVisit(t *testing.T) {
	_, na := NormalCase()
	type args struct {
		from    int
		adj     map[int][]int
		visited map[int]bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal case",
			args: args{
				from: 0,
				adj: na,
				visited: map[int]bool{0: false, 1: false, 2: false, 3: false},
			},
			want: []int {2, 3, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := visit(tt.args.from, tt.args.adj, tt.args.visited); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("visit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasCircuit(t *testing.T) {
	ns, na := NormalCase()
	cs, ca := CircuitCase()
	ccs, cca := ComplexCase()

	type args struct {
		sorted []int
		adj    map[int][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal case",
			args: args{
				sorted: ns,
				adj: na,			
			},
			want: false,
		},
		{
			name: "circuit case",
			args: args{
				sorted: cs,
				adj: ca,
			},
			want: true,
		},
		{
			name: "complex case",
			args: args{
				sorted: ccs,
				adj: cca,
			},
			want: false,
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

