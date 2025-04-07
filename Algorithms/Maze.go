package Algo

import "fmt"

//maze solving algo

func walk(maze [][]string, wall string, curr, end []int, seen [][]bool, path [][]int) bool {
	// fmt.Printf("called walk with : curr:%v end:%v path:%v\n", curr, end, path)
	dir := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	if curr[0] < 0 || curr[1] >= len(maze[0]) ||
		curr[1] < 0 || curr[0] >= len(maze) {
		fmt.Printf("curr[0]:%v<0 || curr[1]:%v >=lenmaze[0]:%v\n", curr[0], curr[1], len(maze[0]))
		fmt.Printf("curr[1]:%v<0 || curr[0]:%v >=lenmaze:%v\n", curr[1], curr[0], len(maze))
		return false
	}
	if maze[curr[0]][curr[1]] == wall {
		fmt.Printf("inside wall curr[0]:%v curr[1]:%v\n", curr[0], curr[0])
		return false
	}
	if end[0] == curr[0] && end[1] == curr[1] {
		fmt.Printf(" end !!!!!! should return true curr[0]:%v curr[1]:%v\n", curr[0], curr[1])
		path = append(path, end)
		fmt.Printf("\npath:%v\n", path)
		return true
	}
	if seen[curr[0]][curr[1]] {
		fmt.Printf("inside seen curr[0]:%v curr[1]:%v\n", curr[0], curr[1])
		return false
	}

	//push
	seen[curr[0]][curr[1]] = true
	path = append(path, curr)

	//recurse
	for i := range len(dir) {
		d := dir[i]
		ncurr := []int{0, 0}
		ncurr[0] = curr[0] + d[0]
		ncurr[1] = curr[1] + d[1]
		fmt.Printf("curr:%v ncurr:%v\n", curr, ncurr)
		if walk(maze, wall, ncurr, end, seen, path) {
			return true
		}
		fmt.Println(" ")
	}

	//pop
	path = path[:len(path)-1]

	return false
}

func SolveMaze(maze [][]string, wall string, start, end []int) (path [][]int) {
	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}
	if walk(maze, wall, start, end, seen, path) {
		fmt.Printf("path:%v", path)
		return path
	}

	return nil
}
