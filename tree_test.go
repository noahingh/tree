package tree

import (
	"reflect"
	"sync"
	"testing"
)

type item string

func (i item) String() string {
	return string(i)
}

func (i item) Less(comp Item) bool {
	return string(i) < string(comp.(item))
}

// getAdj convert the list into the metrix.
func getAdj(l map[int][]int) [CntItemLimit][CntItemLimit]bool {
	var (
		ret [CntItemLimit][CntItemLimit]bool
	)

	for from, tos := range l {
		for _, to := range tos {
			ret[from][to] = true
		}
	}
	return ret
}

func TestRendering(t *testing.T) {
	var (
		root  = item("root")
		dir0  = item("dir 0")
		dir1  = item("dir 1")
		file0 = item("file 0")
		file1 = item("file 1")
		file2 = item("file 2")
	)

	tree := NewTree(root)

	tree.Move(dir0, root)
	tree.Move(file1, dir0)
	tree.Move(file0, dir0)

	tree.Move(dir1, root)
	tree.Move(file2, dir1)

	// move dir1 under dir0.
	tree.Move(dir1, dir0)

	// remove dir 1
	tree.Remove(dir1)

	tt := struct {
		name    string
		tree    *Tree
		want    []string
		wantErr bool
	}{
		name: "rendering",
		tree: tree,
		// childs would be sorted before rendering.
		want: []string{
			"root",
			"└── dir 0",
			"    ├── file 0",
			"    └── file 1",
		},
		wantErr: false,
	}
	t.Run(tt.name, func(t *testing.T) {
		tr := tt.tree
		got, err := tr.Render()
		if (err != nil) != tt.wantErr {
			t.Errorf("Tree.Render() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Tree.Render() = %v, want %v", got, tt.want)
		}
	})
}

func TestTree_Render(t *testing.T) {
	type fields struct {
		root    Item
		adj     [CntItemLimit][CntItemLimit]bool
		entries []Item
		mux     *sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name: "rendering file system",
			fields: fields{
				root: item("root"),
				adj: getAdj(map[int][]int{
					0: []int{1, 4},
					1: []int{2, 3},
					4: []int{5},
				}),
				entries: []Item{
					item("root"),    // 0
					item("dir 0"),   // 1
					item("child 0"), // 2
					item("child 1"), // 3
					item("dir 1"),   // 4
					item("child 2"), // 5
				},
				mux: &sync.Mutex{},
			},
			want: []string{
				"root",
				"├── dir 0",
				"│   ├── child 0",
				"│   └── child 1",
				"└── dir 1",
				"    └── child 2",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tree{
				root:    tt.fields.root,
				adj:     tt.fields.adj,
				entries: tt.fields.entries,
				mux:     tt.fields.mux,
			}
			got, err := tr.Render()
			if (err != nil) != tt.wantErr {
				t.Errorf("Tree.Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.Render() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_compact(t *testing.T) {
	type fields struct {
		root    Item
		adj     [CntItemLimit][CntItemLimit]bool
		entries []Item
		mux     *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want fields
	}{
		{
			name: "simple compact",
			fields: fields {
				root: item("root"),
				adj: getAdj(map[int][]int{
					0: []int {2},
				}),
				entries: []Item{
					item("root"),
					nil,
					item("dir 1"),
					nil,
				},
			},
			want: fields {
				root: item("root"),
				adj: getAdj(map[int][]int{
					0: []int {1},
				}),
				entries: []Item{
					item("root"),
					item("dir 1"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tree{
				root:    tt.fields.root,
				adj:     tt.fields.adj,
				entries: tt.fields.entries,
				mux:     tt.fields.mux,
			}
			tr.compact()

			if !reflect.DeepEqual(tr.adj, tt.want.adj) {
				t.Errorf("Tree.compact(), adj = %v, want %v", tr.adj, tt.want.adj)
			}

			if !reflect.DeepEqual(tr.entries, tt.want.entries) {
				t.Errorf("Tree.compact(), entries = %v, want %v", tr.entries, tt.want.entries)
			}
		})
	}
}

func TestTree_getSubTreeItems(t *testing.T) {
	type fields struct {
		root    Item
		adj     [CntItemLimit][CntItemLimit]bool
		entries []Item
		mux     *sync.Mutex
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "sub-tree items of 'dir 0'",
			fields: fields{
				root: item("root"),
				adj: getAdj(map[int][]int{
					0: []int{1, 4},
					1: []int{2, 3},
					4: []int{5},
				}),
				entries: []Item{
					item("root"),    // 0
					item("dir 0"),   // 1
					item("child 0"), // 2
					item("child 1"), // 3
					item("dir 1"),   // 4
					item("child 2"), // 5
				},
				mux: &sync.Mutex{},
			},
			args: args{
				i: 1,
			},
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tree{
				root:    tt.fields.root,
				adj:     tt.fields.adj,
				entries: tt.fields.entries,
				mux:     tt.fields.mux,
			}
			if got := tr.getSubTreeItems(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.getSubTreeItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_getAdj(t *testing.T) {
	type fields struct {
		root    Item
		adj     [CntItemLimit][CntItemLimit]bool
		entries []Item
		mux     *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]bool
	}{
		{
			name: "converting into the slice",
			fields: fields{
				root: item("root"),
				adj: getAdj(map[int][]int{
					0: []int{1, 2},
					1: []int{3},
				}),
				entries: []Item{
					item("root"),
					item("dir 0"),
					item("dir 1"),
					item("fie 0"),
				},
				mux: &sync.Mutex{},
			},
			want: [][]bool{
				{false, true, true, false},
				{false, false, false, true},
				{false, false, false, false},
				{false, false, false, false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tree{
				root:    tt.fields.root,
				adj:     tt.fields.adj,
				entries: tt.fields.entries,
				mux:     tt.fields.mux,
			}
			if got := tr.getAdj(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.getAdj() = %v, want %v", got, tt.want)
			}
		})
	}
}
