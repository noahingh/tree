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

func TestRendering(t *testing.T) {
	var (
		root = item("root")
		dir0 = item("dir 0")
		dir1 = item("dir 1")
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

	tt := struct{
		name string
		tree *Tree
		want []string
		wantErr bool
	}{
		name: "rendering",
		tree: tree,
		// childs would be sorted before rendering.
		want: []string{
			"root",
			"└── dir 0",
			"    ├── dir 1",		
			"    │   └── file 2",
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
			name: "normal",
			fields: fields{
				root: item("root"),
				adj: [CntItemLimit][CntItemLimit]bool{
					{false, true, false, false, true, false},
					{false, false, true, true, false, false},
					{false, false, false, false, false, false},
					{false, false, false, false, false, false},
					{false, false, false, false, false, true},
					{false, false, false, false, false, false},
				},
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

