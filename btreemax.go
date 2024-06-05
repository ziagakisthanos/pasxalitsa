package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	current := root
	for root.Right != nil {
		root = BTreeMax(root.Right)
		//  root = root.Right
	}
	if root.Data > current.Data {
		current = root
	}
	// for root.Left != nil {
	// 	root = BTreeMax(root.Left)
	//  root = root.Left
	// }
	return current
}
