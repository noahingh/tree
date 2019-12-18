package tree

import (
	"fmt"
)

var (
	// ErrCircuit is the error when the topology has a circuit.
	ErrCircuit = fmt.Errorf("Circuit exist in the topology")
)

// HasCircuit validate the set of adjacents has a circuit or not.
func HasCircuit(adj [][]bool) bool {
	_, err := topologicalSort(adj)
	if err != nil {
		return true
	}
	return false
}

func topologicalSort(adj [][]bool) ([]int, error) {
	var (
		lenVer  = len(adj)
		sorted  = []int{}
		visited = make([]bool, lenVer)
	)

	if lenVer == 0 {
		return []int{}, nil
	}

	if !isSquare(adj) {
		return []int{}, fmt.Errorf("the adj must to be square")
	}

	for from := 0; from < len(adj); from++ {
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
func visit(from int, adj [][]bool, visited []bool) []int {
	var (
		output []int
		lenVer = len(adj)
	)

	// set visited
	visited[from] = true

	for to := 0; to < lenVer; to++ {
		if adj[from][to] && !visited[to] {
			o := visit(to, adj, visited)
			output = append(output, o...)
		}
	}

	output = append(output, from)
	return output
}

func isSquare(adj [][]bool) bool {
	l := len(adj)

	for _, s := range adj {
		if len(s) != l {
			return false
		}
	}
	return true
}

func hasCircuit(reversed []int, adj [][]bool) bool {
	for right := len(reversed) - 1; right >= 1; right-- {
		for left := right - 1; left >= 0; left-- {
			from, to := reversed[right], reversed[left]

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
