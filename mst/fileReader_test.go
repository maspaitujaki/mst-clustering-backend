package mst

type addTestNode struct {
	fileName      string
	expectedNodes []Node
}

var addTestNodes = []addTestNode{
	addTestNode{"test/tc1.txt", []Node{
		Node{"1", 6.0, 8.0},
		Node{"2", 7.0, 4.0},
		Node{"3", 8.0, 3.0},
		Node{"4", 2.0, 1.0},
		Node{"5", 1.0, 2.0},
	}},
}

// func TestReadNodes(t *testing.T) {
// 	for _, test := range addTestNodes {
// 		nodes, err := ReadNodes(test.fileName)
// 		if err != nil {
// 			t.Errorf("ReadNodes(%s) returned error: %v", test.fileName, err)
// 		}
// 		if len(nodes) != len(test.expectedNodes) {
// 			t.Errorf("ReadNodes(%s) returned wrong number of nodes: %d", test.fileName, len(nodes))
// 		}
// 		for i, node := range nodes {
// 			if !IsNodeEqual(node, test.expectedNodes[i]) {
// 				t.Errorf("ReadNodes(%s) returned wrong node: %v", test.fileName, node)
// 			}
// 		}
// 	}
// }
