package main

// TreeNode /**
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/binary-tree-preorder-traversal/
/**
 * 先序 中-左-右
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

//https://leetcode.cn/problems/binary-tree-inorder-traversal/
/**
 * 中序 左-中-右
 * 1.将结点的左子树全部依次压栈
 * 2.当左子树到底，依次弹出结点打印，对弹出结点的右结点压栈，重复步骤1
 * 3.没有子树且栈空 结束
 */
func inorderTraversal(root *TreeNode) []int {
	treeNodeArr := make([]int, 0)
	treeNodeStack := make([]*TreeNode, 0)
	if root != nil {
		for len(treeNodeStack) > 0 || root != nil {
			if root != nil {
				treeNodeStack = append(treeNodeStack, root)
				root = root.Left
			} else {
				root = treeNodeStack[len(treeNodeStack)-1]
				treeNodeStack = treeNodeStack[:len(treeNodeStack)-1]

				treeNodeArr = append(treeNodeArr, root.Val)
				root = root.Right
			}
		}
	}

	return treeNodeArr
}

//https://leetcode.cn/problems/binary-tree-postorder-traversal/description/
/**
 * 后序遍历 左-右-中
 * 0. 2个栈实现
 * 1. 可以参考先序遍历: 当前结点先入s1栈，如果s1非空，弹出结点入s2栈
 * 2. s1栈依次入当前节点左，右结点 循环1步骤
 * 3. 使得入s2栈的节点顺序为中-右-左
 * 4. 依次弹出s2栈节点 即为左-右-中
 */
func postorderTraversal(root *TreeNode) []int {
	stack1 := make([]*TreeNode, 0)
	stack2 := make([]*TreeNode, 0)
	treeNodeArr := make([]int, 0)

	if root != nil {
		stack1 = append(stack1, root)
		for len(stack1) > 0 {
			node := stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
			stack2 = append(stack2, node)
			if node.Left != nil {
				stack1 = append(stack1, node.Left)
			}

			if node.Right != nil {
				stack1 = append(stack1, node.Right)
			}
		}
	}

	if len(stack2) > 0 {
		for idx := len(stack2) - 1; idx >= 0; idx-- {
			treeNodeArr = append(treeNodeArr, stack2[idx].Val)
		}
	}

	return treeNodeArr
}
