package leetcode

/*
  一些恶魔抓住了公主（P）并将她关在了地下城的右下角。地下城是由 M x N 个房间组成的二维网格。我们英勇的骑士（K）最初被安置在左上角的房间里，他必须穿过地下城并通过对抗恶魔来拯救公主。

  骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至 0 或以下，他会立即死亡。

  有些房间由恶魔守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；其他房间要么是空的（房间里的值为 0），要么包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。

  为了尽快到达公主，骑士决定每次只向右或向下移动一步。

  说明:
    骑士的健康点数没有上限。
    任何房间都可能对骑士的健康点数造成威胁，也可能增加骑士的健康点数，包括骑士进入的左上角房间以及公主被监禁的右下角房间。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/dungeon-game
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 0
	}
	dp := make([]int, len(dungeon[0]))
	for i := len(dungeon) - 1; i >= 0; i-- {
		for j := len(dungeon[0]) - 1; j >= 0; j-- {
			tmp := 1<<31 - 1
			if i < len(dungeon)-1 {
				tmp = Min(tmp, dp[j])
			}
			if j < len(dungeon[0])-1 {
				tmp = Min(tmp, dp[j+1])
			}
			if tmp == 1<<31-1 {
				tmp = 1
			}
			if tmp-dungeon[i][j] > 0 {
				dp[j] = tmp - dungeon[i][j]
			} else {
				dp[j] = 1
			}
		}
	}
	return dp[0]
}
