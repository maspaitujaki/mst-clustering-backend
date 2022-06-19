package mst

import "math"

type Node struct {
	Name string
	X, Y float64
}

func IsNodeEqual(a, b Node) bool {
	return a.Name == b.Name && a.X == b.X && a.Y == b.Y
}

func Distance(a, b Node) float64 {
	return math.Sqrt(math.Pow(b.X-a.X, 2) + math.Pow(b.Y-a.Y, 2))
}
