package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == elem {
		return root
	}

	if elem < root.Data {
		if BTreeSearchItem(root.Left, elem) != nil {
			BTreeSearchItem(root.Left, elem).Parent = root
			return BTreeSearchItem(root.Left, elem)
		}
	}
	if BTreeSearchItem(root.Right, elem) != nil {
		BTreeSearchItem(root.Right, elem).Parent = root
		return BTreeSearchItem(root.Right, elem)
	}
	return nil
}
