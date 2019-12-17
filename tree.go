package tree

import (
	"fmt"
	"sync"
	"sort"
)

const (
	// CntItemLimit is the limit count of nodes.
	CntItemLimit = 1000
)

var (
	// ErrCntItemLimit is the error of the limit count.
	ErrCntItemLimit = fmt.Errorf("the count of items exceed the limit")
	// ErrItemNotExist is the error if the item doesn't exist in the tree.
	ErrItemNotExist = fmt.Errorf("the item doesn't exist")
	// ErrParentNotExist is the error if the parent doesn't exist in the tree.
	ErrParentNotExist = fmt.Errorf("the parent doesn't exist")
)

// NewTree return the tree which set the root with the item.
func NewTree(item Item) *Tree {
	t := &Tree{
		root: item,
	}

	t.addEntry(item)
	return t
}

// Tree is directory structure to return "tree" command format.
type Tree struct {
	root    Item
	adj     [CntItemLimit][CntItemLimit]bool
	entries []Item

	mux sync.Mutex
}

// Render return the tree format strings but if there is a circuit in the tree it return the error.
func (t *Tree) Render() ([]string, error) {
	// validate a circuit exist or not.

	root := t.root
	return t.render(root), nil
}

func (t *Tree) render(i Item) []string {
	var (
		ret []string
	)
	ret = append(ret, i.String())

	childs := t.getChilds(i)
	
	if len(childs) == 0 {
		return ret
	}

	for i, c := range childs {
		lines := t.render(c)

		if i == len(childs)-1 {
			lines = renderLastChild(lines)
		} else {
			lines = renderChild(lines)
		}

		for _, l := range lines {
			ret = append(ret, l)
		}
	}
	return ret
}

func (t *Tree) getChilds(i Item) Items {
	var (
		childs Items
		idxi, _ = t.getIndexEqual(i)
	)

	for idxc := 0; idxc < len(t.entries); idxc++ {
		if t.adj[idxi][idxc] {
			child := t.entries[idxc]
			childs = append(childs, child)
		}
	}

	sort.Sort(childs)

	return childs
}

// Move move the item into the parent.
func (t *Tree) Move(item, parent Item) error {
	t.mux.Lock()
	defer t.mux.Unlock()

	if !t.has(parent) {
		return ErrParentNotExist
	}

	if !t.has(item) {
		t.addEntry(item)
	}

	if err := t.move(item, parent); err != nil {
		return err
	}
	return nil
}

func (t *Tree) move(child, parent Item) error {
	idxc, _ := t.getIndexEqual(child)
	idxp, _ := t.getIndexEqual(parent)

	idxOld, err := t.getIndexParent(child)

	if err != nil {
		t.adj[idxp][idxc] = true
		return nil
	}

	t.adj[idxOld][idxc] = false
	t.adj[idxp][idxc] = true

	return nil
}

func (t *Tree) has(i Item) bool {
	if _, err := t.getIndexEqual(i); err != nil {
		return false
	}
	return true
}

func (t *Tree) getIndexParent(child Item) (int, error) {
	idxc, _ := t.getIndexEqual(child)

	for idxp := 0; idxp < len(t.entries); idxp++ {
		if t.adj[idxp][idxc] {
			return idxp, nil
		}
	}

	return -1, ErrParentNotExist
}

func (t *Tree) getIndexEqual(i Item) (int, error) {
	for idx, comp := range t.entries {
		// equal
		if !i.Less(comp) && !comp.Less(i) {
			return idx, nil
		}
	}
	return -1, ErrItemNotExist
}

func (t *Tree) addEntry(i Item) int {
	t.entries = append(t.entries, i)
	return len(t.entries)
}
