package main

// Tags:
// Union-Find
func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	uf := NewUnionFind(m * n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				continue
			}
			x := i*n + j
			if i+1 < m && grid[i+1][j] == 0 {
				y := (i+1)*n + j
				uf.Union(x, y)
			}
			if j+1 < n && grid[i][j+1] == 0 {
				y := i*n + j + 1
				uf.Union(x, y)
			}
		}
	}

	valid := make([]bool, m*n)
	for i := 0; i < m*n; i++ {
		r, c := i/n, i%n
		if grid[r][c] == 1 {
			continue
		}
		valid[uf.Find(i)] = true
	}

	for i := 0; i < m*n; i++ {
		x := uf.Find(i)
		if !valid[x] {
			continue
		}
		r, c := i/n, i%n
		if r == 0 || r == m-1 || c == 0 || c == n-1 {
			valid[x] = false
		}
	}

	out := 0
	for i := 0; i < m*n; i++ {
		if valid[i] {
			out++
		}
	}
	return out
}

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return UnionFind{
		parent: parent,
		size:   size,
	}
}

func (uf UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf UnionFind) Union(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}
	if uf.size[x] < uf.size[y] {
		x, y = y, x
	}
	uf.parent[y] = x
	uf.size[x] += uf.size[y]
}

// BFS + DFS
func closedIsland2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	inValid := make([][]bool, m)
	for i := range inValid {
		inValid[i] = make([]bool, n)
	}

	q := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && grid[i][j] == 0 {
				q = append(q, []int{i, j})
			}
		}
	}

	for len(q) > 0 {
		i, j := q[0][0], q[0][1]
		q = q[1:]
		inValid[i][j] = true
		for _, f := range forwards {
			ni, nj := i+f[0], j+f[1]
			if ni < 0 || m-1 < ni || nj < 0 || n-1 < nj || grid[ni][nj] == 1 || inValid[ni][nj] {
				continue
			}
			q = append(q, []int{ni, nj})
		}
	}

	out := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 || inValid[i][j] {
				continue
			}
			out++
			dfs(grid, m, n, inValid, i, j)
		}
	}
	return out
}

var forwards = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func dfs(grid [][]int, m, n int, inValid [][]bool, i, j int) {
	inValid[i][j] = true
	for _, f := range forwards {
		ni, nj := i+f[0], j+f[1]
		if ni < 0 || m-1 < ni || nj < 0 || n-1 < nj || grid[ni][nj] == 1 || inValid[ni][nj] {
			continue
		}
		dfs(grid, m, n, inValid, ni, nj)
	}
}
