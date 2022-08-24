package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	paretns := node.Parent

	rplc.Parent = paretns

	if paretns.Data > node.Data {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}
	return root
}
