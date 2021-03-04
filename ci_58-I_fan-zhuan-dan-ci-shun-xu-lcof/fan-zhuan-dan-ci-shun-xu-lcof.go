package main

import "strings"

/*
  输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。

  示例 1：
    输入: "the sky is blue"
    输出: "blue is sky the"

  示例 2：
    输入: "  hello world!  "
    输出: "world! hello"
    解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

  示例 3：
    输入: "a good   example"
    输出: "example good a"
    解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

  说明：
    1. 无空格字符构成一个单词。
    2. 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
    3. 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reverseWords(s string) string {
	strs := []string{}
	for s != "" {
		for s != "" && s[0] == ' ' {
			s = s[1:]
		}
		if s == "" {
			break
		}
		i := 0
		for i < len(s) && s[i] != ' ' {
			i++
		}
		strs = append(strs, s[:i])
		s = s[i:]

	}
	for i := 0; i < len(strs)>>1; i++ {
		strs[i], strs[len(strs)-1-i] = strs[len(strs)-1-i], strs[i]
	}
	return strings.Join(strs, " ")
}
