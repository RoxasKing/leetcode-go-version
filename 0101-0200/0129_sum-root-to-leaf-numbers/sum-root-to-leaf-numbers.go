package main

/*
  You are given the root of a binary tree containing digits from 0 to 9 only.

  Each root-to-leaf path in the tree represents a number.
    For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
  Return the total sum of all root-to-leaf numbers.

  A leaf node is a node with no children.

  Example 1:
    Input: root = [1,2,3]
    Output: 25
    Explanation:
    The root-to-leaf path 1->2 represents the number 12.
    The root-to-leaf path 1->3 represents the number 13.
    Therefore, sum = 12 + 13 = 25.

  Example 2:
    Input: root = [4,9,0,5,1]
    Output: 1026
    Explanation:
    The root-to-leaf path 4->9->5 represents the number 495.
    The root-to-leaf path 4->9->1 represents the number 491.
    The root-to-leaf path 4->0 represents the number 40.
    Therefore, sum = 495 + 491 + 40 = 1026.

  Constraints:
    1. The number of nodes in the tree is in the range [1, 1000].
    2. 0 <= Node.val <= 9
    3. The depth of the tree will not exceed 10.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sum-root-to-leaf-numbers
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	out := 0
	dfs(root, root.Val, &out)
	return out
}

func dfs(root *TreeNode, cur int, out *int) {
	if root.Left == nil && root.Right == nil {
		*out += cur
	}

	if root.Left != nil {
		dfs(root.Left, cur*10+root.Left.Val, out)
	}

	if root.Right != nil {
		dfs(root.Right, cur*10+root.Right.Val, out)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
