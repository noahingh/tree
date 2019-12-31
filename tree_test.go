package tree

import (
	"sync"
	"testing"
)

func TestTree_Move(t *testing.T) {
	type fields struct {
		root    *node
		entries nodes
		mux     *sync.Mutex
	}
	type args struct {
		item   Item
		parent Item
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tree{
				root:    tt.fields.root,
				entries: tt.fields.entries,
				mux:     tt.fields.mux,
			}
			if err := tr.Move(tt.args.item, tt.args.parent); (err != nil) != tt.wantErr {
				t.Errorf("Tree.Move() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_move(t *testing.T) {
	type args struct {
		child  *node
		parent *node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "no parent",
			args: args{
				child: &node{item: item("new")},
				parent: &node{item: item("root")},
			},
		},
		{
			name: "has parent",
			args: args{
				child: &node{
					parent: &node{
						item: item("prev"),
						children: nodes{ 
							&node{
								item: item("child"),
							},
						},
					},
					item: item("child"),
				},
				parent: &node{item: item("root")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			move(tt.args.child, tt.args.parent)
		})
	}
}
