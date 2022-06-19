package clustering

import (
	"backend/mst"
	"math"
)

func MakeCluster(t mst.Tree, n int) [][]mst.Node {
	if len(t.Nodes) == 0 {
		return nil
	}
	if len(t.Nodes) < n {
		return nil
	}
	t.Edges = mst.QuickSortEdgesDesc(t.Edges)
	// for _, e := range t.Edges {
	// 	fmt.Printf("From %s to %s cost %f\n", e.From.Name, e.To.Name, e.Weight)
	// }
	trees := make([]mst.Tree, 0)
	trees = append(trees, t)

	for i := n; i > 1; i-- {
		var popped mst.Edge
		currenTree := trees[getMaxTree(trees)]
		trees = mst.RemoveIndex(trees, getMaxTree(trees))
		currenTree.Edges, popped = popEdges(currenTree.Edges)
		newTree1 := mst.Tree{Nodes: []mst.Node{popped.From}, Edges: []mst.Edge{}}
		newTree2 := mst.Tree{Nodes: []mst.Node{popped.To}, Edges: []mst.Edge{}}

		j := 0
		for len(currenTree.Edges) > 0 {
			if j >= len(currenTree.Edges) {
				j = 0
			}
			edge := currenTree.Edges[j]
			if nodesContain(newTree1.Nodes, edge.From) {
				newTree1.Edges = append(newTree1.Edges, edge)
				newTree1.Nodes = append(newTree1.Nodes, edge.To)
				currenTree.Edges = removeIndex(currenTree.Edges, j)
				continue
			}
			if nodesContain(newTree1.Nodes, edge.To) {
				newTree1.Edges = append(newTree1.Edges, edge)
				newTree1.Nodes = append(newTree1.Nodes, edge.From)
				currenTree.Edges = removeIndex(currenTree.Edges, j)
				continue
			}

			if nodesContain(newTree2.Nodes, edge.From) {
				newTree2.Edges = append(newTree2.Edges, edge)
				newTree2.Nodes = append(newTree2.Nodes, edge.To)
				currenTree.Edges = removeIndex(currenTree.Edges, j)
				continue
			}
			if nodesContain(newTree2.Nodes, edge.To) {
				newTree2.Edges = append(newTree2.Edges, edge)
				newTree2.Nodes = append(newTree2.Nodes, edge.From)
				currenTree.Edges = removeIndex(currenTree.Edges, j)
				continue
			}
			j++

		}
		newTree1.Edges = mst.QuickSortEdgesDesc(newTree1.Edges)
		newTree2.Edges = mst.QuickSortEdgesDesc(newTree2.Edges)
		trees = append(trees, newTree1)
		trees = append(trees, newTree2)
	}
	result := make([][]mst.Node, 0)
	for _, tree := range trees {
		result = append(result, tree.Nodes)
	}
	return result
}

func popEdges(e []mst.Edge) ([]mst.Edge, mst.Edge) {
	r := e[0]
	e = append(e[:0], e[0+1:]...)
	return e, r
}

func removeIndex(t []mst.Edge, index int) []mst.Edge {
	return append(t[:index], t[index+1:]...)
}

func getMaxTree(t []mst.Tree) int {
	max := math.Inf(-1)
	maxIdx := -1
	for i, tree := range t {
		if len(tree.Edges) == 0 {
			continue
		}
		if tree.Edges[0].Weight > max {
			max = tree.Edges[0].Weight
			maxIdx = i
		}
	}
	return maxIdx
}

func nodesContain(nodes []mst.Node, node mst.Node) bool {
	for _, nodei := range nodes {
		if mst.IsNodeEqual(nodei, node) {
			return true
		}
	}
	return false
}
