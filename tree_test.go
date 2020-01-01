package tree

import (
	"reflect"
	"testing"
)

func TestIntegration(t *testing.T) {
	tree := NewTree(item("root"))

	tree.Move(item("dir0"), item("root"))
	tree.Move(item("dir1"), item("root"))

	tree.Move(item("file0"), item("dir0"))
	tree.Move(item("file1"), item("dir0"))

	tree.Move(item("file2"), item("dir1"))
	tree.Move(item("dir1-0"), item("dir1"))
	t.Run("add items", func(t *testing.T) {
		expected := []string{
			"root",
			"├── dir0",
			"│   ├── file0",
			"│   └── file1",
			"└── dir1",
			"    ├── dir1-0",
			"    └── file2",
		}
		if got := tree.Render(); !reflect.DeepEqual(got, expected) {
			t.Errorf("tree.Render() = %s, want %s", got, expected)
		}
	})

	tree.Move(item("dir0"), item("dir1"))
	t.Run("move dir0", func(t *testing.T) {
		expected := []string{
			"root",
			"└── dir1",
			"    ├── dir0",
			"    │   ├── file0",
			"    │   └── file1",
			"    ├── dir1-0",
			"    └── file2",
		}
		if got := tree.Render(); !reflect.DeepEqual(got, expected) {
			t.Errorf("tree.Render() = %s, want %s", got, expected)
		}
	})

	tree.Remove(item("dir0"))
	t.Run("remove dir0", func(t *testing.T) {
		expected := []string{
			"root",
			"└── dir1",
			"    ├── dir1-0",
			"    └── file2",
		}
		if got := tree.Render(); !reflect.DeepEqual(got, expected) {
			t.Errorf("tree.Render() = %s, want %s", got, expected)
		}
	})
}
