package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	// Helper: Find minimum node in subtree rooted at n
	minimum := func(n *TreeNode) *TreeNode {
		curr := n
		for curr.Left != nil {
			curr = curr.Left
		}
		return curr
	}

	// Case 1: node has no left child -> transplant right child
	if node.Left == nil {
		root = BTreeTransplant(root, node, node.Right)
	} else if node.Right == nil {
		// Case 2: node has no right child -> transplant left child
		root = BTreeTransplant(root, node, node.Left)
	} else {
		// Case 3: node has two children
		// Find successor (min in right subtree)
		successor := minimum(node.Right)
		if successor.Parent != node {
			// Replace successor by its right child
			root = BTreeTransplant(root, successor, successor.Right)
			// Link successor's right child to node's right child
			successor.Right = node.Right
			if successor.Right != nil {
				successor.Right.Parent = successor
			}
		}
		// Replace node by successor
		root = BTreeTransplant(root, node, successor)
		// Link successor's left child to node's left child
		successor.Left = node.Left
		if successor.Left != nil {
			successor.Left.Parent = successor
		}
	}
	return root
}
