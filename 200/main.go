package main

type pair struct {
	i, j int
}

type que []pair

func (q *que) pop() pair {
	tmp := (*q)[0]
	if len((*q)) == 1 {
		(*q) = (*q)[0:0]
		return tmp
	}
	(*q) = (*q)[1:]
	return tmp
}

func (q *que) add(p pair) {
	(*q) = append((*q), p)
}

func bfs(adj [][]byte, start pair, visited *map[pair]struct{}, maxRow, maxCol int) {
	q := make(que, 0)
	(*visited)[start] = struct{}{}
	q.add(start)
	for len(q) != 0 {
		firstIn := q.pop()
		if _, ok := (*visited)[pair{firstIn.i - 1, firstIn.j}]; !ok && firstIn.i-1 >= 0 && adj[firstIn.i-1][firstIn.j] == '1' {
			q.add(pair{firstIn.i - 1, firstIn.j})
			(*visited)[pair{firstIn.i - 1, firstIn.j}] = struct{}{}
		}
		if _, ok := (*visited)[pair{firstIn.i, firstIn.j - 1}]; !ok && firstIn.j-1 >= 0 && adj[firstIn.i][firstIn.j-1] == '1' {
			q.add(pair{firstIn.i, firstIn.j - 1})
			(*visited)[pair{firstIn.i, firstIn.j - 1}] = struct{}{}
		}
		if _, ok := (*visited)[pair{firstIn.i + 1, firstIn.j}]; !ok && firstIn.i+1 < maxRow && adj[firstIn.i+1][firstIn.j] == '1' {
			q.add(pair{firstIn.i + 1, firstIn.j})
			(*visited)[pair{firstIn.i + 1, firstIn.j}] = struct{}{}
		}
		if _, ok := (*visited)[pair{firstIn.i, firstIn.j + 1}]; !ok && firstIn.j+1 < maxCol && adj[firstIn.i][firstIn.j+1] == '1' {
			q.add(pair{firstIn.i, firstIn.j + 1})
			(*visited)[pair{firstIn.i, firstIn.j + 1}] = struct{}{}
		}
	}
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	visited := make(map[pair]struct{}, 0)
	countIslands := 0
	for r := 0; r < len(grid); r++ {
		for l := 0; l < len(grid[0]); l++ {
			if _, ok := visited[pair{r, l}]; grid[r][l] == '1' && !ok {
				bfs(grid, pair{r, l}, &visited, len(grid), len(grid[0]))
				countIslands += 1
			}
		}
	}
	return countIslands
}
