package tree

import (
	"fmt"
	"sync"
	"sort"
)

// NewTree return the tree which set the root with the item.
func NewTree(item Item) *Tree {
	root := newNode(item)
	entries := nodes{
		root,
	}

	t := &Tree{
		root:    root,
		entries: entries,
		mux:     &sync.Mutex{},
	}

	return t
}

// Tree is directory structure to return "tree" command format.
type Tree struct {
	root    *node
	entries nodes
	mux     *sync.Mutex
}

// Render return the tree format strings.
func (t *Tree) Render() []string {
	t.mux.Lock()
	defer t.mux.Unlock()

	return render(t.root)
}

// Has check if the item is included or not
func (t *Tree) Has(item Item) bool {
	t.mux.Lock()
	defer t.mux.Unlock()
	return t.has(item)
}

func (t *Tree) has(i Item) bool {
	if n := t.search(i); n != nil {
		return true
	}
	return false
}

func (t *Tree) search(i Item) *node {
	tmp := newNode(i)

	idx := sort.Search(len(t.entries), func (i int) bool {
		n := t.entries[i]
		if tmp.equal(n) {
			return true
		}

		return false
	})

	if idx == len(t.entries) {
		return nil
	}

	return t.entries[idx]
}

// Move move the item into the parent.
func (t *Tree) Move(item, parent Item) error {
	t.mux.Lock()
	defer t.mux.Unlock()

	var c, p *node
	
	if parent != nil && !t.has(item) {
		return fmt.Errorf("the parent item doesn't exist")
	}

    if parent == nil {
		p = t.root
	} else {
		p = t.search(parent)
	}

	if !t.has(item) {
		c = newNode(item)
		t.move(c, p)
		return nil
	} 

	c = t.search(item)
	// if c is the ancestor of p, it can't move.
	if isAncestor(c, p) {
		return fmt.Errorf("couldn't move the item")
	}
	t.move(c, p)

	return nil
}

func (t *Tree) move(child, parent *node) {
	prev := child.parent
	prev.children = removeNode(prev.children, child)

	parent.children = appendNode(parent.children, child)
	child.parent = parent
}

// Remove remove the item in the tree, and also all childs under the item.
func (t *Tree) Remove(item Item) error {
	t.mux.Lock()
	defer t.mux.Unlock()

	if !t.has(item) {
		return fmt.Errorf("the item is not exist")
	}

	n := t.search(item)

	p := n.parent
	p.children = removeNode(p.children, n)

	n.parent = nil
	t.remove(n)

	return nil
}

func (t *Tree) remove(n *node) {
	if len(n.children) == 0 {
		removeNode(t.entries, n)
		n.parent = nil
		return 
	}

	for _, c := range n.children {
		t.remove(c)
	}
	removeNode(t.entries, n)
	n.children = nodes{}
}
