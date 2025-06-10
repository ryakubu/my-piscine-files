package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	// If node is the root, rplc becomes the new root
	if node.Parent == nil {
		if rplc != nil {
			rplc.Parent = nil
		}
		return rplc
	}

	// Update parent's child pointer to point to rplc instead of node
	if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}

	// Update rplc's parent pointer if rplc is not nil
	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}
