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

After creating a new tree, you can construct the tree by adding or removing the item by methods, Move and Remove:

	tree.Move(file("dir 0"), file("root"))
	tree.Move(file("file 0"), file("dir 0"))
	
	tree.Remove(file("dir 0"))

Note that when you construct the tree you should consider the circuit, of course, 
this package prevent the tree has a circuit when moving a item. 

At last, you can render the tree by Render method: 

	ss := tree.Render()
	for _, s := range ss {
		fmt.println(s)
	}

*/
package tree
