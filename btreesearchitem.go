package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == elem {
		return root
	}
	n := BTreeSearchItem(root.Left, elem)

	if n == nil {
		n = BTreeSearchItem(root.Right, elem)
	}

	return n
}
