package tree

import (
	"strings"
	"testing"
)

func newItem(n string) *item {
	return &item{
		name: n,
	}
}

type item struct {
	name string
}

func (i *item) Less(comp Item) bool {
	return i.name < comp.(*item).name
}

func (i *item) String() string {
	return i.name
}

func TestRender(t *testing.T) {
	expected := `root
└── dir
    ├── child
    └── lastChild`

	root := newItem("root")
	tree := NewTree(root)

	dir := newItem("dir")
	tree.Move(dir, root)

	child := newItem("child")
	tree.Move(child, dir)

	lastChild := newItem("lastChild")
	tree.Move(lastChild, dir)

	r, _ := tree.Render()
	if output := strings.Join(r, "\n"); output != expected {
		t.Errorf("output is not equal to expected.\n %s", output)
	}
}
