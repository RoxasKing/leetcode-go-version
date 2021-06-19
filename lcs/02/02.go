package main

import "sort"

/*
  有 N 位扣友参加了微软与力扣举办了「以扣会友」线下活动。主办方提供了 2*N 道题目，整型数组 questions 中每个数字对应了每道题目所涉及的知识点类型。
  若每位扣友选择不同的一题，请返回被选的 N 道题目至少包含多少种知识点类型。

  示例 1：
    输入：questions = [2,1,6,2]
    输出：1
    解释：有 2 位扣友在 4 道题目中选择 2 题。
         可选择完成知识点类型为 2 的题目时，此时仅一种知识点类型
         因此至少包含 1 种知识点类型。

  示例 2：
    输入：questions = [1,5,1,3,4,5,2,5,3,3,8,6]
    输出：2
    解释：有 6 位扣友在 12 道题目中选择题目，需要选择 6 题。
         选择完成知识点类型为 3、5 的题目，因此至少包含 2 种知识点类型。

  提示：
    1. questions.length == 2*n
    2. 2 <= questions.length <= 10^5
    3. 1 <= questions[i] <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/WqXACV
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Greedy Algorithm

func halfQuestions(questions []int) int {
	n := len(questions)
	cnt := map[int]int{}
	for _, q := range questions {
		cnt[q]++
	}
	arr := make([]int, 0, len(cnt))
	for num := range cnt {
		arr = append(arr, num)
	}
	sort.Slice(arr, func(i, j int) bool { return cnt[arr[i]] > cnt[arr[j]] })
	out := 0
	sum := 0
	for _, num := range arr {
		out++
		sum += cnt[num]
		if sum >= n>>1 {
			break
		}
	}
	return out
}
