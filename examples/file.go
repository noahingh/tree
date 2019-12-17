package main

import (
	"fmt"

	"github.com/hanjunlee/tree"
)

func newFile(n string) *file {
	return &file{
		name: n,
	}
}

type file struct {
	name string
}

func (f *file) String() string {
	return f.name
}

func (f *file) Less(comp tree.Item) bool {
	if f.name < comp.(*file).name {
		return true
	}

	return false
}

func main() {

	root := newFile("root")
	t := tree.NewTree(root)

	dir0 := newFile("dir 0")
	dir1 := newFile("dir 1")

	c0 := newFile("child 0")
	c1 := newFile("child 1")
	c2 := newFile("child 2")

	t.Move(dir0, root)
	t.Move(dir1, root)

	t.Move(c0, dir0)
	t.Move(c1, dir0)

	t.Move(c2, dir1)

	result, _ := t.Render()
	for _, l := range  result {
		fmt.Println(l)
	}
}
