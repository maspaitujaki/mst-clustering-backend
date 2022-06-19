package mst

import "math/rand"

type Edge struct {
	From, To Node
	Weight   float64
}

func QuickSortEdgesAsc(a []Edge) []Edge {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i].Weight < a[right].Weight {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuickSortEdgesAsc(a[:left])
	QuickSortEdgesAsc(a[left+1:])

	return a
}

func QuickSortEdgesDesc(a []Edge) []Edge {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i].Weight > a[right].Weight {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuickSortEdgesDesc(a[:left])
	QuickSortEdgesDesc(a[left+1:])

	return a
}
