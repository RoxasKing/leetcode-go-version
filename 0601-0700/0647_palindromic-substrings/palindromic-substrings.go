package main

import "strings"

func countSubstrings(s string) int {
	out := 0
	n := len(s)
	for i := 0; i < n; i++ {
		out += countPalindrome(s, n, i, i)
		if i+1 < n && s[i] == s[i+1] {
			out += countPalindrome(s, n, i, i+1)
		}
	}
	return out
}

func countPalindrome(s string, n, l, r int) int {
	out := 1
	for ; l-1 >= 0 && r+1 < n && s[l-1] == s[r+1]; l, r = l-1, r+1 {
		out++
	}
	return out
}

// Dynamic Programming
func countSubstrings2(s string) int {
	n := len(s)
	f := make([]bool, n)

	out := 0
	for i := 0; i < n; i++ {
		f[i] = true
		out++
		for j := 0; j < i; j++ {
			if s[i] == s[j] && f[j+1] {
				f[j] = true
				out++
			} else {
				f[j] = false
			}
		}
	}
	return out
}

// Manacher Algorithm
func countSubstrings3(s string) int {
	ss := strings.Split(s, "")
	t := strings.Join(ss, "#")
	t = "$#" + t + "#!"
	n := len(t) - 1

	f := make([]int, n)
	iMax, rMax, out := 0, 0, 0
	for i := 1; i < n; i++ {
		if i < rMax {
			f[i] = Min(rMax-i+1, f[iMax-(i-iMax)])
		} else {
			f[i] = 1
		}

		for t[i+f[i]] == t[i-f[i]] {
			f[i]++
		}
		if i+f[i]-1 > rMax {
			iMax = i
			rMax = i + f[i] - 1
		}
		out += f[i] / 2
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
