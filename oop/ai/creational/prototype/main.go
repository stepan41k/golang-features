package main

type Node struct {
	Name string
	Children []string
}

func (n *Node) Clone() *Node {
	newChildren := make([]string, len(n.Children))

	copy(newChildren, n.Children)

	return &Node{
		Name: n.Name,
		Children: newChildren,
	}
}