package main

/*
  给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。

  示例:
    输入: [1,2,3,4,5]
    输出: [120,60,40,30,24]

  提示：
    1. 所有元素乘积之和不会溢出 32 位整数
    2. a.length <= 100000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func constructArr(a []int) []int {
	n := len(a)
	if n == 0 {
		return nil
	}
	out := make([]int, n)
	out[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		out[i] = out[i+1] * a[i+1]
	}
	mul := 1
	for i := 0; i < n; i++ {
		out[i] *= mul
		mul *= a[i]
	}
	return out
}
