package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	size := BTreeLevelCount(root)
	for i := 0; i < size; i++ {
		Helper(root, i, f)
	}
}

func Helper(root *TreeNode, i int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if i == 0 {
		f(root.Data)
	} else {
		Helper(root.Left, i-1, f)
		Helper(root.Right, i-1, f)
	}
}
