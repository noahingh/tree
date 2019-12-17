package tree

import (
	"fmt"
	"sync"
)

const (
	tempNode = -1
)

// NewTree return the tree which set the root with the item.
func NewTree(item Item) *Tree {
	t := &Tree{
		nextID: int(0),
	}
	root := t.addEntry(item)
	t.root = root

	return t
}

// Tree is directory structure to return "tree" command format.
type Tree struct {
	nextID  int
	root    *node
	entries nodes

	mux sync.Mutex
}

// Render return the result of which the root is rendered as strings.
func (t *Tree) Render() ([]string, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	adj := t.getAdjacents()
	// sort adjactents with the topological sorting to validate the circuit exist or not.
	if _, err := topologicalSort(adj); err != nil {
		return []string{}, fmt.Errorf("circuit exist in the tree")
	}

	return t.root.Render(), nil
}

// Move move the item into the parent. 
func (t *Tree) Move(item, parent Item) error {
	t.mux.Lock()
	defer t.mux.Unlock()

	p := t.getEqual(parent)
	if p == nil {
		return fmt.Errorf("'%s' parent node doesn't exist", p)
	}

	if same := t.getEqual(item); same != nil {
		return fmt.Errorf("'%s' node already exist", same)
	}

	n := t.addEntry(item)
	return t.move(n, p)
}

func (t *Tree) move(node, parent *node) error {
	parent.AddChild(node)

	return nil
}

func (t *Tree) getEqual(i Item) *node {
	n := newNode(tempNode, i)
	for _, comp := range t.entries {
		// equal
		if !n.Less(comp) && !comp.Less(n) {
			return comp
		}
	}

	return nil
}

func (t *Tree) addEntry(i Item) *node {
	n := newNode(t.nextID, i)
	t.nextID++

	t.entries = append(t.entries, n)
	return n
}

// getAdjacents return the set of edges between nodes.
func (t *Tree) getAdjacents() map[int][]int {
	adj := make(map[int][]int, 0)

	for _, from := range t.entries {
		childs := make([]int, 0)

		for _, to := range from.childs {
			childs = append(childs, to.id)
		}
		adj[from.id] = childs
	}
	return adj
}
