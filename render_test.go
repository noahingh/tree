package tree

import (
	"reflect"
	"testing"
)

func Test_render(t *testing.T) {
	type args struct {
		n *node
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "file system",
			args: args{
				n: &node{
					item: item("root"),
					children: nodes{
						&node{
							item: item("dir 0"),
							children: nodes{
								&node{
									item: item("file 0"),
									children: nodes{},
								},
								&node{
									item: item("file 1"),
									children: nodes{},
								},
							},
						},
					},
				},
			},
			want: []string{
				"root",
				"└── dir 0",
				"    ├── file 0",
				"    └── file 1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := render(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("render() = %v, want %v", got, tt.want)
			}
		})
	}
}
