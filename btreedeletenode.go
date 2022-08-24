package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == node {
		if root.Left != nil {
			min := BTreeMin(root.Left)
		}
	}

	if node.Data > root.Data {
		return BTreeDeleteNode(root.Right, node)
	} else {
		return BTreeDeleteNode(root.Left, node)
	}
}
