package piscine

func BTreeLevelCount(root *TreeNode) int {
	count := 1
	if root == nil {
		return count
	}
	count = SizeOfOneN(root, count)
	return count
}

func SizeOfOneN(root *TreeNode, count int) int {
	if root == nil {
		return 0
	}
	r := SizeOfOneN(root.Right, count+1)
	l := SizeOfOneN(root.Left, count+1)
	if r >= l && r > count {
		return r
	} else if l > r && l > count {
		return l
	} else {
		return count
	}
}
