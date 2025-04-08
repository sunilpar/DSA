package Algo

// maze solving algo
func walk(maze [][]string, wall string, curr, end []int, seen [][]bool, path [][]int) ([][]int, bool) {
	dir := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	if curr[0] < 0 || curr[1] >= len(maze[0]) ||
		curr[1] < 0 || curr[0] >= len(maze) {
		return path, false
	}
	if maze[curr[0]][curr[1]] == wall {
		return path, false
	}
	if end[0] == curr[0] && end[1] == curr[1] {
		path = append(path, end)
		return path, true
	}
	if seen[curr[0]][curr[1]] {
		return path, false
	}

	// push
	seen[curr[0]][curr[1]] = true
	path = append(path, curr)

	// recurse
	for _, d := range dir {
		ncurr := []int{curr[0] + d[0], curr[1] + d[1]}
		updatedPath, found := walk(maze, wall, ncurr, end, seen, path)
		if found {
			return updatedPath, true
		}
	}

	// pop
	path = path[:len(path)-1]
	return path, false
}

func SolveMaze(maze [][]string, wall string, start, end []int) (path [][]int) {
	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}
	finalPath, found := walk(maze, wall, start, end, seen, path)
	if found {
		return finalPath
	}
	return nil
}
