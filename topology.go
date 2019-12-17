package tree

import (
	"fmt"
)

var (
	// ErrCircuit is the error when the topology has a circuit.
	ErrCircuit = fmt.Errorf("Circuit exist in the topology")
)

// HasCircuit validate the set of adjacents has a circuit or not.
func HasCircuit(adj [CntItemLimit][CntItemLimit]bool) bool {
	_, err := topologicalSort(adj)
	if err != nil {
		return true
	}
	return false
}

func topologicalSort(adj [CntItemLimit][CntItemLimit]bool) ([]int, error) {
	var (
		sorted  = []int{}
		visited = make([]bool, CntItemLimit)
	)

	for from := 0; from < CntItemLimit; from++ {
		if visited[from] {
			continue
		}

		output := visit(from, adj, visited)
		sorted = append(sorted, output...)
	}

	sorted = reverse(sorted)

	if hasCircuit(sorted, adj) {
		return []int{}, ErrCircuit
	}

	return sorted, nil
}

// visit visit all adjacent vertex as the depth first search.
func visit(from int, adj [CntItemLimit][CntItemLimit]bool, visited []bool) []int {
	var (
		output []int
	)

	// set visited
	visited[from] = true

	for to := 0; to < CntItemLimit; to++ {
		if adj[from][to] && !visited[to] {
			o := visit(to, adj, visited)
			output = append(output, o...)
		}
	}

	output = append(output, from)
	return output
}

func hasCircuit(sorted []int, adj [CntItemLimit][CntItemLimit]bool) bool {
	for right := len(sorted) - 1; right >= 1; right-- {
		for left := right - 1; left >= 0; left-- {
			from, to := sorted[right], sorted[left]

			if adj[from][to] {
				return true
			}
		}
	}
	return false
}

func reverse(ints []int) (ret []int) {
	for i := len(ints) - 1; i >= 0; i-- {
		ret = append(ret, ints[i])
	}
	return ret
}
