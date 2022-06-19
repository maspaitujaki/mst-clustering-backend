package mst

import (
	"encoding/csv"
	"os"
	"strconv"
)

func ReadNodes(fileName string) (string, string, []Node, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", "", nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	row, err := csvReader.Read()
	if err != nil {
		panic(err)
	}
	xAtr := row[1]
	yAtr := row[2]

	rows, err := csvReader.ReadAll() // `rows` is of type [][]string
	if err != nil {
		panic(err)
	}

	var nodes []Node
	for _, row := range rows {
		var X, Y float64
		name := row[0]
		if s, err := strconv.ParseFloat(row[1], 64); err == nil {
			X = s
		}
		if s, err := strconv.ParseFloat(row[2], 64); err == nil {
			Y = s
		}
		nodes = append(nodes, Node{name, X, Y})
	}
	return xAtr, yAtr, nodes, nil
}

func ReadEdges(nodes []Node) ([]Edge, error) {
	var edges []Edge
	for i, node := range nodes {
		for j, node2 := range nodes {
			if j > i {
				edges = append(edges, Edge{node, node2, Distance(node, node2)})
			}
		}
	}

	return edges, nil
}
