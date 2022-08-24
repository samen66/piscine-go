package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil || root.Right == nil || root.Left == nil {
		return true
	}

	if root.Data > root.Left.Data && root.Data <= root.Right.Data {
		return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
	} else {
		return false
	}
}
