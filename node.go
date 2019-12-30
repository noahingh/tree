package tree

import "sort"

type nodes []*node

func (ns nodes) Len() int {
	return len(ns)
}

func (ns nodes) Less(i, j int) bool {
	return ns[i].less(ns[j])
}

func (ns nodes) Swap(i, j int) {
	temp := ns[i]
	ns[i] = ns[j]
	ns[j] = temp
	return
}

func (ns nodes) append(n *node) {
	ns = append(ns, n)
	sort.Sort(ns)
	return
}

type node struct {
	parent   *node
	children nodes

	item Item
}

func newNode(i Item) *node {
	return &node{
		item: i,
	}
}

func (n *node) equal(comp *node) bool {
	if !n.less(comp) && !comp.less(n) {
		return true
	}
	return false
}

func (n *node) less(comp *node) bool {
	return n.item.Less(comp.item)
}

func (n *node) String() string {
	return n.item.String()
}

func isAncestor(parent, child *node) bool {
	if child.parent == nil {
		return false
	}

	if child.parent.equal(parent) {
		return true
	}

	return isAncestor(parent, child.parent)
}
