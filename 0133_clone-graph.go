package leetcode

/*
  给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。
  图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。

  class Node {
      public int val;
      public List<Node> neighbors;
  }

  测试用例格式：
  简单起见，每个节点的值都和它的索引相同。例如，第一个节点值为 1（val = 1），第二个节点值为 2（val = 2），以此类推。该图在测试用例中使用邻接列表表示。
  邻接列表 是用于表示有限图的无序列表的集合。每个列表都描述了图中节点的邻居集。
  给定节点将始终是图中的第一个节点（值为 1）。你必须将 给定节点的拷贝 作为对克隆图的引用返回。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/clone-graph
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	newN := &Node{Val: node.Val}
	queue := []*Node{node}
	clone := []*Node{newN}
	mark := make(map[int]*Node)
	mark[newN.Val] = newN
	for len(queue) != 0 {
		srcN, dstN := queue[0], clone[0]
		queue, clone = queue[1:], clone[1:]
		if len(dstN.Neighbors) != 0 {
			continue
		}
		for _, srcn := range srcN.Neighbors {
			var dstn *Node
			if n, ok := mark[srcn.Val]; ok {
				dstn = n
			} else {
				dstn = &Node{Val: srcn.Val}
				mark[dstn.Val] = dstn
			}
			dstN.Neighbors = append(dstN.Neighbors, dstn)
			queue = append(queue, srcn)
			clone = append(clone, dstn)
		}
	}
	return newN
}

type Node struct {
	Val       int
	Neighbors []*Node
}
