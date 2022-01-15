package treeNode

// Node N叉树结构
type Node struct {
	Val      int
	Children []*Node
}

// NewNode 构造一颗N叉树
func NewNode() *Node {
	/**
			 	1
		3    	2      4
	5      6
	*/
	root := &Node{Val: 4}
	root.Children = append(root.Children, &Node{Val: 3})
	root.Children = append(root.Children, &Node{Val: 2})
	root.Children = append(root.Children, &Node{Val: 4})

	root.Children[0].Children = append(root.Children[0].Children, &Node{Val: 5})
	root.Children[0].Children = append(root.Children[0].Children, &Node{Val: 6})
	return root
}
