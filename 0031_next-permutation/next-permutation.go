package main

/*
  实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
  如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
  必须原地修改，只允许使用额外常数空间。

  以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
    1,2,3 → 1,3,2
    3,2,1 → 1,2,3
    1,1,5 → 1,5,1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/next-permutation
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	reverse(nums[i+1:])
	if i < 0 {
		return
	}
	j := bSearch(nums, nums[i], i+1, n-1)
	nums[i], nums[j] = nums[j], nums[i]
}

func reverse(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func bSearch(nums []int, target, l, r int) int {
	for l < r {
		m := l + (r-l)>>1
		if nums[m] <= target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
