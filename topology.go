package tree

import (
	"fmt"
)

var (
	// ErrCircuit is the error when the topology has a circuit.
	ErrCircuit = fmt.Errorf("Circuit exist in the topology")
)

func topologicalSort(adj map[int][]int) ([]int, error) {
	var (
		sorted  = []int{}
		visited = map[int]bool{}
	)

	for from := range adj {
		visited[from] = false
	}

	for from := range adj {
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
func visit(from int, adj map[int][]int, visited map[int]bool) []int {
	var (
		output []int
	)

	// set visited
	visited[from] = true

	tos := adj[from]
	for _, to := range tos {
		if visited[to] {
			continue
		}

		o := visit(to, adj, visited)
		output = append(output, o...)
	}

	output = append(output, from)
	return output
}

func hasCircuit(sorted []int, adj map[int][]int) bool {
	for right := len(sorted) - 1; right >= 1; right-- {
		for left := right - 1; left >= 0; left-- {
			from, to := sorted[right], sorted[left]

			if tos := adj[from]; searchInt(tos, to) {
				return true
			}
		}
	}
	return false
}

func searchInt(ints []int, x int) bool {
	for _, i := range ints {
		if i == x {
			return true
		}
	}
	return false
}

func reverse(ints []int) (ret []int) {
	for i := len(ints)-1; i >= 0; i-- {
		ret = append(ret, ints[i])
	}
	return ret
}
