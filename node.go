package tree

import (
	"sort"
	"sync"
)

var (
	// TabChild is the prepended string for childs when a node would be rendered.
	TabChild = "├── "
	// TabGrandChild is the prepended string for grand childs when a node would be rendered.
	TabGrandChild = "│   "
	// TabLastChild is the prepended string for the last child when a node would be rendered.
	TabLastChild = "└── "
	// TabGrandLastChild is the prepended string for the last child when a node would be rendered.
	TabGrandLastChild = "    "
)

type nodes []*node

func (ns nodes) Len() int {
	return len(ns)
}

func (ns nodes) Less(i, j int) bool {
	return ns[i].Less(ns[j])
}

func (ns nodes) Swap(i, j int) {
	tmp := ns[j]
	ns[j] = ns[i]
	ns[i] = tmp
	return
}

func newNode(id int, i Item) *node {
	return &node{
		id: id,
		item: i,
	}
}

type node struct {
	id     int
	item   Item
	childs nodes

	mux sync.Mutex
}

func (n *node) String() string {
	return n.item.String()
}

func (n *node) Less(comp *node) bool {
	return n.item.Less(comp.item)
}

func (n *node) AddChild(c *node) {
	n.mux.Lock()
	defer n.mux.Unlock()

	n.childs = append(n.childs, c)
	sort.Sort(n.childs)
	return
}

// Render return the result of which the node is rendered as strings.
func (n *node) Render() []string {
	var (
		ret []string
	)

	ret = append(ret, n.String())

	if len(n.childs) == 0 {
		return ret
	}

	for i, c := range n.childs {
		lines := c.Render()

		if i == len(n.childs)-1 {
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

func renderChild(lines []string) []string {
	var ret []string

	l, lines := lines[0], lines[1:]

	ret = append(ret, TabChild+l)
	for _, l := range lines {
		ret = append(ret, TabGrandChild+l)
	}
	return ret
}

func renderLastChild(lines []string) []string {
	var ret []string

	l, lines := lines[0], lines[1:]

	ret = append(ret, TabLastChild+l)
	for _, l := range lines {
		ret = append(ret, TabGrandLastChild+l)
	}
	return ret
}
