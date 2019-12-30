package main

import (
	"fmt"

	"github.com/hanjunlee/tree"
)

type file string

func (f file) String() string {
	return string(f)
}

func (f file) Less(comp tree.Item) bool {
	return string(f) < string(comp.(file))
}

func main() {
	t := tree.NewTree(file("root"))

	t.Move(file("dir0"), file("root"))
	t.Move(file("dir1"), file("root"))

	t.Move(file("file 0"), file("dir0"))
	t.Move(file("file 1"), file("dir0"))

	t.Move(file("file 2"), file("dir1"))

	result := t.Render()
	for _, l := range result {
		fmt.Println(l)
	}
}
