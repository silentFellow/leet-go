package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	valMap map[int]struct{}
}

func Constructor(root *TreeNode) FindElements {
	valMap := make(map[int]struct{})
	root = constructTree(root, 0, valMap)
	return FindElements{valMap: valMap}
}

func (this *FindElements) Find(target int) bool {
	if _, ok := this.valMap[target]; ok {
		return true
	}

  return false
}

func constructTree(root *TreeNode, val int, valMap map[int]struct{}) *TreeNode {
	if root == nil {
		return nil
	}

	root.Val = val
	valMap[val] = struct{}{}
	root.Left = constructTree(root.Left, 2*val+1, valMap)
	root.Right = constructTree(root.Right, 2*val+2, valMap)

	return root
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
