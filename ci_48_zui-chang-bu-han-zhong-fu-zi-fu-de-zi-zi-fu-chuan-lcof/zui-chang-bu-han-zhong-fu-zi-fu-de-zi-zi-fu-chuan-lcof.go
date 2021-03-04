package main

/*
  请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。

  示例 1:
    输入: "abcabcbb"
    输出: 3
    解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

  示例 2:
    输入: "bbbbb"
    输出: 1
    解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

  示例 3:
    输入: "pwwkew"
    输出: 3
    解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
         请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

  提示：
    s.length <= 40000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Sliding Window
func lengthOfLongestSubstring(s string) int {
	out := 0
	cnt := [128]int{}
	for l, r := 0, 0; r < len(s); r++ {
		cnt[s[r]]++
		for l < r && cnt[s[r]] > 1 {
			cnt[s[l]]--
			l++
		}
		if r+1-l > out {
			out = r + 1 - l
		}
	}
	return out
}
