package main

import "strings"

/*
  n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

  上图为 8 皇后问题的一种解法。

  给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
  每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

  示例：
    输入：4
    输出：[
     [".Q..",  // 解法 1
      "...Q",
      "Q...",
      "..Q."],

     ["..Q.",  // 解法 2
      "Q...",
      "...Q",
      ".Q.."]
    ]
    解释: 4 皇后问题存在两个不同的解法。

  提示：
    皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/n-queens
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Backtracking
func solveNQueens(n int) [][]string {
	track := make([]string, n)
	for i := range track {
		track[i] = strings.Repeat(".", n)
	}
	cols := make([]bool, n)
	mainDiag := make([]bool, n<<1-1)
	antiDiag := make([]bool, n<<1-1)
	out := [][]string{}

	backtrack(n, track, cols, mainDiag, antiDiag, &out, 0)

	return out
}

func backtrack(n int, track []string, cols, mainDiag, antiDiag []bool, out *[][]string, row int) {
	if row == n {
		tmp := make([]string, n)
		copy(tmp, track)
		*out = append(*out, tmp)
		return
	}

	for col := 0; col < n; col++ {
		if cols[col] || mainDiag[row+col] || antiDiag[row-col+n-1] {
			continue
		}

		bs := []byte(track[row])

		bs[col] = 'Q'
		track[row] = string(bs)
		cols[col] = true
		mainDiag[row+col] = true
		antiDiag[row-col+n-1] = true

		backtrack(n, track, cols, mainDiag, antiDiag, out, row+1)

		bs[col] = '.'
		track[row] = string(bs)
		cols[col] = false
		mainDiag[row+col] = false
		antiDiag[row-col+n-1] = false
	}
}
