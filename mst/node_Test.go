package mst

import "testing"

func TestDistance(t *testing.T) {
	a := Node{"a", 1.0, 2.0}
	b := Node{"b", 3.0, 4.0}
	if Distance(a, b) != 5.0 {
		t.Errorf("Distance(%v, %v) returned %f", a, b, Distance(a, b))
	}
}
