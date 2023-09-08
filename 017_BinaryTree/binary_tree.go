package main

// TreeNode /**
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/binary-tree-preorder-traversal/
/**
 * 先序 中左右
 * 非递归 先压头结点
 * 当栈非空，当前结点出栈 先压右结点 再压左结点
 */
func preorderTraversal(root *TreeNode) []int {
	treeNodeArr := make([]int, 0)
	treeNodeStack := make([]*TreeNode, 0)
	if root != nil {
		treeNodeStack = append(treeNodeStack, root)
		for len(treeNodeStack) > 0 {
			node := treeNodeStack[len(treeNodeStack)-1]
			treeNodeStack = treeNodeStack[:len(treeNodeStack)-1]
			treeNodeArr = append(treeNodeArr, node.Val)

			if node.Right != nil {
				treeNodeStack = append(treeNodeStack, node.Right)
			}

			if node.Left != nil {
				treeNodeStack = append(treeNodeStack, node.Left)
			}
		}
	}

	return treeNodeArr
}
