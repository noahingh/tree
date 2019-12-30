package tree

import (
	"reflect"
	"testing"
)

func Test_appendNode(t *testing.T) {
	type args struct {
		ns   nodes
		news []*node
	}
	tests := []struct {
		name string
		args args
		want nodes
	}{
		// TODO: Add test cases.
		{
			name: "append nodes",
			args: args{
				ns: nodes{
					&node{
						item: item("item 0"),
					},
				},
				news: []*node{
					&node{
						item: item("item 1"),
					},
					&node{
						item: item("item 2"),
					},
				},
			},
			want: nodes{
				&node{
					item: item("item 0"),
				},
				&node{
					item: item("item 1"),
				},
				&node{
					item: item("item 2"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendNode(tt.args.ns, tt.args.news...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeNode(t *testing.T) {
	type args struct {
		ns nodes
		n  *node
	}
	tests := []struct {
		name string
		args args
		want nodes
	}{
		// TODO: Add test cases.
		{
			name: "remove node",
			args: args{
				ns: nodes{
					&node{
						item: item("item 0"),
					},
					&node{
						item: item("item 1"),
					},
				},
				n: &node{
					item: item("item 1"),
				},
			},
			want: nodes{
				&node{
					item: item("item 0"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNode(tt.args.ns, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
