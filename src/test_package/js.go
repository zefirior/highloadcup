package test_package

type Example struct {
	Row   string
	Nodes []Node
}

type Node struct {
	Left  string
	Right string
}

func (n *Node) Equal(other *Node) bool {
	return n.Left == other.Left && n.Right == other.Right
}

func (e *Example) Equal(other *Example) bool {
	if e.Row != other.Row {
		return false
	}

	n := len(e.Nodes)
	if n != len(other.Nodes) {
		return false
	}

	for i := 0; i < n; i++ {
		if !e.Nodes[i].Equal(&other.Nodes[i]) {
			return false
		}
	}
	return true
}
