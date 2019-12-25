/*
Package tree help to display content recursively as tree-like format. 

Before create a new tree, implement Item:

	type file string

	func (f file) String() string {
		return string(f)
	}

	func (f file) Less(comp Item) bool {
		return string(f) < string(comp.(file))
	}

Firstly create a new tree with the item: 

	tree := NewTree(file("root"))

After creating a new tree, you can construct the tree by adding or removing the node by methods, Move and Remove:

	tree.Move(file("dir 0"), file("root"))
	tree.Move(file("file 0"), file("dir 0"))
	
	tree.Remove(file("dir 0"))

Note that when you construct the tree you should consider the circuit of the tree for a infinite loop, of course, 
this package validates that the circuit exist or not before removing and rendering. These methods return the error so that 
you should check the error.

At last, you can render the tree by Render method: 

	ss := tree.Render()
	for _, s := range ss {
		fmt.println(s)
	}

IMPORTANT: At this moment, this package has the limit of count of nodes, it is 1000, if you add more than the limit it return the error.
*/
package tree
