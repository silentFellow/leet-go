package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	n := len(preorder)

	hmap := make(map[int]int)
	for i, v := range postorder {
		hmap[v] = i
	}

	var build func(i1, i2, j1 int) *TreeNode
	build = func(i1, i2, j1 int) *TreeNode {
		if i1 > i2 {
			return nil
		}

		root := &TreeNode{Val: preorder[i1]}
		if i1 != i2 {
			leftVal := preorder[i1+1]

			m := hmap[leftVal]
			leftSize := m - j1 + 1

			root.Left = build(i1+1, i1+leftSize, j1)
			root.Right = build(i1+leftSize+1, i2, m+1)
		}

		return root
	}

	return build(0, n-1, 0)
}
