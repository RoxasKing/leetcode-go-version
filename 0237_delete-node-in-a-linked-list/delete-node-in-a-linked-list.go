package main

/*
  请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点。传入函数的唯一参数为 要被删除的节点 。

  现有一个链表 -- head = [4,5,1,9]，它可以表示为:

  提示：
    链表至少包含两个节点。
    链表中所有节点的值都是唯一的。
    给定的节点为非末尾节点并且一定是链表中的一个有效节点。
    不要从你的函数中返回任何结果。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/delete-node-in-a-linked-list
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func deleteNode(node *ListNode) {
	cur := node
	pre := &ListNode{Next: cur}
	for cur.Next != nil {
		cur.Val = cur.Next.Val
		cur, pre = cur.Next, pre.Next
	}
	pre.Next = nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
