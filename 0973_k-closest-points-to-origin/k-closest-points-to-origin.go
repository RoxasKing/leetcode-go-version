package main

import (
	"container/heap"
	"math/rand"
)

/*
  我们有一个由平面上的点组成的列表 points。需要从中找出 K 个距离原点 (0, 0) 最近的点。
  （这里，平面上两点之间的距离是欧几里德距离。）
  你可以按任何顺序返回答案。除了点坐标的顺序之外，答案确保是唯一的。

  示例 1：
    输入：points = [[1,3],[-2,2]], K = 1
    输出：[[-2,2]]
    解释：
    (1, 3) 和原点之间的距离为 sqrt(10)，
    (-2, 2) 和原点之间的距离为 sqrt(8)，
    由于 sqrt(8) < sqrt(10)，(-2, 2) 离原点更近。
    我们只需要距离原点最近的 K = 1 个点，所以答案就是 [[-2,2]]。

  示例 2：
    输入：points = [[3,3],[5,-1],[-2,4]], K = 2
    输出：[[3,3],[-2,4]]
    （答案 [[-2,4],[3,3]] 也会被接受。）

  提示：
    1 <= K <= points.length <= 10000
    -10000 < points[i][0] < 10000
    -10000 < points[i][1] < 10000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/k-closest-points-to-origin
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Quick Sort
func kClosest(points [][]int, K int) [][]int {
	quickSort(points, 0, len(points)-1, K)
	return points[:K]
}

func quickSort(points [][]int, l, r int, k int) {
	if l == r {
		return
	}
	pivotIndex := l + rand.Intn(r+1-l)
	points[pivotIndex], points[r] = points[r], points[pivotIndex]
	index := l
	for i := l; i < r; i++ {
		if isCloser(points[i], points[r]) {
			points[i], points[index] = points[index], points[i]
			index++
		}
	}
	points[index], points[r] = points[r], points[index]
	if index < k-1 {
		quickSort(points, index+1, r, k)
	} else if index > k-1 {
		quickSort(points, l, index-1, k)
	}
}

func isCloser(a, b []int) bool { return a[0]*a[0]+a[1]*a[1] < b[0]*b[0]+b[1]*b[1] }

// Priority Queue(Heap Sort)
func kClosest2(points [][]int, K int) [][]int {
	n := len(points)
	pq := make(PriorityQueue, n)
	for i := range pq {
		pq[i] = point{
			val:  points[i][0]*points[i][0] + points[i][1]*points[i][1],
			pair: points[i],
		}
	}
	heap.Init(&pq)
	out := make([][]int, 0, K)
	for len(out) < K {
		out = append(out, heap.Pop(&pq).(point).pair)
	}
	return out
}

type point struct {
	val  int
	pair []int
}
type PriorityQueue []point

func (p PriorityQueue) Len() int            { return len(p) }
func (p PriorityQueue) Less(i, j int) bool  { return p[i].val < p[j].val }
func (p PriorityQueue) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *PriorityQueue) Push(x interface{}) { *p = append(*p, x.(point)) }
func (p *PriorityQueue) Pop() interface{} {
	last := len(*p) - 1
	out := (*p)[last]
	*p = (*p)[:last]
	return out
}
