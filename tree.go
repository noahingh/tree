package tree

import (
	"fmt"
	"sort"
	"sync"
)

const (
	// CntItemLimit is the limit count of nodes.
	CntItemLimit = 1000

	notExist = -1
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
		mux:  &sync.Mutex{},
	}

	t.addEntry(item)
	return t
}

// Tree is directory structure to return "tree" command format.
type Tree struct {
	root Item
	// The index of entry should be the id of the item,
	// all methods of the Tree use the index instead of the Item.
	adj     [CntItemLimit][CntItemLimit]bool
	entries []Item

	mux *sync.Mutex
}

// Render return the tree format strings but if there is a circuit in the tree it return the error.
func (t *Tree) Render() ([]string, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	// validate a circuit exist or not.
	if HasCircuit(t.getAdj()) {
		return nil, ErrCircuit
	}

	r := getIndexEqual(t.entries, t.root)
	return t.render(r), nil
}

func (t *Tree) render(i int) []string {
	var (
		ret  []string
		item = t.entries[i]
	)
	ret = append(ret, item.String())

	childs := t.getChilds(i)

	if len(childs) == 0 {
		return ret
	}

	for i, c := range childs {
		lines := t.render(c)

		if i == len(childs)-1 {
			lines = tabLastChild(lines)
		} else {
			lines = tabChild(lines)
		}

		for _, l := range lines {
			ret = append(ret, l)
		}
	}
	return ret
}

func (t *Tree) getChilds(i int) []int {
	var (
		childItems = make(Items, 0)
		childs     = make([]int, 0)
	)

	for child := 0; child < len(t.entries); child++ {
		if t.adj[i][child] {
			childItems = append(childItems, t.entries[child])
		}
	}

	// sort childs to print ordered.
	sort.Sort(childItems)

	for _, item := range childItems {
		childs = append(childs, getIndexEqual(t.entries, item))
	}

	return childs
}

// Move move the item into the parent.
func (t *Tree) Move(item, parent Item) error {
	t.mux.Lock()
	defer t.mux.Unlock()

	var (
		i = getIndexEqual(t.entries, item)
		p = getIndexEqual(t.entries, parent)
	)

	if p == notExist {
		return ErrParentNotExist
	}

	if i == notExist {
		i = t.addEntry(item)
	}

	if err := t.move(i, p); err != nil {
		return err
	}
	return nil
}

func (t *Tree) move(child, parent int) error {
	oldParent := t.getParent(child)

	if oldParent == notExist {
		t.adj[parent][child] = true
		return nil
	}

	t.adj[oldParent][child] = false
	t.adj[parent][child] = true

	return nil
}

// Remove remove the item in the tree, and also all childs under the item.
func (t *Tree) Remove(item Item) error {
	t.mux.Lock()
	defer t.mux.Unlock()

	i := getIndexEqual(t.entries, item)

	items := t.getSubTreeItems(i)
	if hasCircuit(items, t.getAdj()) {
		return ErrCircuit
	}

	for _, i := range items {
		t.remove(i)
	}

	t.compact()
	return nil
}

func (t *Tree) remove(i int) {
	// remove the item from internal data.
	t.entries[i] = nil

	for from := 0; from < len(t.entries); from++ {
		t.adj[from][i] = false
	}

	for to := 0; to < len(t.entries); to++ {
		t.adj[i][to] = false
	}
}

// compact remove the nil and shift the element to empty.
func (t *Tree) compact() {
	var (
		prevID  = make(map[int]int)
		entries = make([]Item, 0)
		adj     = [CntItemLimit][CntItemLimit]bool{}
	)

	// compact the entries of the tree.
	for old, i := range t.entries {
		if i == nil {
			continue
		}

		entries = append(entries, i)
		new := len(entries) - 1
		prevID[new] = old
	}

	// set new adjacents.
	for from := 0; from < len(entries); from++ {
		for to := 0; to < len(entries); to++ {
			prevFrom, prevTo := prevID[from], prevID[to]

			if t.adj[prevFrom][prevTo] {
				adj[from][to] = true
			}
		}
	}

	t.entries = entries
	t.adj = adj
}

//  getSubTreeItems return childs and grand childs of the index by the slice which is sorted topologically.
func (t *Tree) getSubTreeItems(i int) []int {
	var (
		visited = make(map[int]bool)
	)

	ret := t.visit(i, visited)
	ret = reverse(ret)

	return ret
}

// depth first search
func (t *Tree) visit(i int, visited map[int]bool) []int {
	var (
		ret = make([]int, 0)
	)

	visited[i] = true

	for child := 0; child < len(t.entries); child++ {
		if _, ok := visited[child]; !ok && t.adj[i][child] {
			grandChilds := t.visit(child, visited)
			ret = append(ret, grandChilds...)
		}
	}

	ret = append(ret, i)
	return ret
}

func (t *Tree) getAdj() [][]bool {
	var (
		size = len(t.entries)
		ret  = make([][]bool, size)
	)

	for i := 0; i < size; i++ {
		ret[i] = make([]bool, size)
	}

	// copy values of adj
	for from := 0; from < size; from++ {
		for to := 0; to < size; to++ {
			if t.adj[from][to] {
				ret[from][to] = true
			}
		}
	}
	return ret
}

func (t *Tree) getParent(child int) int {
	for parent := 0; parent < len(t.entries); parent++ {
		if t.adj[parent][child] {
			return parent
		}
	}

	return -1
}

func (t *Tree) addEntry(i Item) int {
	t.entries = append(t.entries, i)
	return len(t.entries) - 1
}

// getIndexEqual return the index, if the item doesn't exist it returns -1.
func getIndexEqual(items []Item, i Item) int {
	for idx, comp := range items {
		// equal
		if !i.Less(comp) && !comp.Less(i) {
			return idx
		}
	}
	return -1
}
