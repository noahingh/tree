package tree

import (
	"sort"
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

func render(n *node) []string {
	var (
		ret  = []string{n.String()}
		cntC = len(n.children)
	)
	if cntC == 0 {
		return ret
	}

	sort.Sort(n.children)
	for i, c := range n.children {
		lines := render(c)

		if i == cntC-1 {
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

func tabChild(lines []string) []string {
	var ret []string

	l, lines := lines[0], lines[1:]

	ret = append(ret, TabChild+l)
	for _, l := range lines {
		ret = append(ret, TabGrandChild+l)
	}
	return ret
}

func tabLastChild(lines []string) []string {
	var ret []string

	l, lines := lines[0], lines[1:]

	ret = append(ret, TabLastChild+l)
	for _, l := range lines {
		ret = append(ret, TabGrandLastChild+l)
	}
	return ret
}
