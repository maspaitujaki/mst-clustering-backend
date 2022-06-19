package mst

func KruskalMST(nodes []Node, edges []Edge) Tree {
	var trees []Tree
	tempNode := make([]Node, 0)
	tempNode = append(tempNode, nodes...)
	count := 0
	for _, edge := range edges {
		count += 1
		// fmt.Println("len Trees: ", len(trees))
		// fmt.Println("len TempNode: ", len(tempNode))
		fromIdx := treesIdxContain(trees, edge.From)
		toIdx := treesIdxContain(trees, edge.To)
		if len(trees) == 1 && len(tempNode) == 0 {
			break
		}
		if fromIdx == -1 && toIdx == -1 {
			tree := Tree{Nodes: []Node{edge.From, edge.To}, Edges: []Edge{edge}}
			trees = append(trees, tree)
			tempNode = RemoveElement(tempNode, edge.From)
			tempNode = RemoveElement(tempNode, edge.To)
			continue
		}
		if fromIdx == -1 && toIdx != -1 {
			trees[toIdx].Nodes = append(trees[toIdx].Nodes, edge.From)
			trees[toIdx].Edges = append(trees[toIdx].Edges, edge)
			tempNode = RemoveElement(tempNode, edge.From)
			continue
		}
		if fromIdx != -1 && toIdx == -1 {
			trees[fromIdx].Nodes = append(trees[fromIdx].Nodes, edge.To)
			trees[fromIdx].Edges = append(trees[fromIdx].Edges, edge)
			tempNode = RemoveElement(tempNode, edge.To)
			continue
		}
		if fromIdx != -1 && toIdx != -1 {
			if fromIdx == toIdx {
				continue
			}
			// fmt.Printf("From %s to %s cost %f\n", edge.From.Name, edge.To.Name, edge.Weight)
			// fmt.Printf("FromIdx %d, ToIdx %d\n", fromIdx, toIdx)
			treeFrom := trees[fromIdx]
			treeTo := trees[toIdx]
			if fromIdx > toIdx {
				trees = RemoveIndex(trees, fromIdx)
				trees = RemoveIndex(trees, toIdx)
			} else {
				trees = RemoveIndex(trees, toIdx)
				trees = RemoveIndex(trees, fromIdx)
			}
			newTree := combineTrees(treeFrom, treeTo)
			newTree.Edges = append(newTree.Edges, edge)
			trees = append(trees, newTree)
			continue
		}
	}
	// fmt.Println(count)
	if len(trees) == 1 {
		return trees[0]
	}
	return Tree{}
}

func treesIdxContain(trees []Tree, node Node) int {
	for i, t := range trees {
		for _, n := range t.Nodes {
			if n.Name == node.Name {
				return i
			}
		}
	}
	return -1
}

func combineTrees(t1, t2 Tree) Tree {
	var nodes []Node
	var edges []Edge
	nodes = append(nodes, t1.Nodes...)
	nodes = append(nodes, t2.Nodes...)
	edges = append(edges, t1.Edges...)
	edges = append(edges, t2.Edges...)
	return Tree{Nodes: nodes, Edges: edges}
}

func RemoveIndex(t []Tree, index int) []Tree {
	return append(t[:index], t[index+1:]...)
}

func RemoveElement(t []Node, n Node) []Node {
	for i, node1 := range t {
		if node1.Name == n.Name {
			return append(t[:i], t[i+1:]...)
		}
	}
	return t
}
