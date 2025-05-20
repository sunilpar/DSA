package leetcode

// import "fmt"
//
//	func GameOfLife(board [][]int) {
//		org := []int{}
//		for i, _ := range board {
//			for j, _ := range board[i] {
//				org = heart(i, j, board)
//				if board[i][j] == 0 {
//					if org[0] > 3 {
//						board[i][j] = 1
//					}
//				} else {
//					if org[1] < 2 || org[1] > 3 {
//						board[i][j] = 0
//					}
//				}
//			}
//		}
//
//		fmt.Printf("Output:%+v\n", board)
//	}
//
//	func heart(i, j int, board [][]int) []int {
//		org := []int{0, 0}
//		O := 0
//		I := 0
//		indexes := [][]int{}
//		indexes = append(indexes, []int{i - 1, j})     //up
//		indexes = append(indexes, []int{i - 1, j - 1}) //up left
//		indexes = append(indexes, []int{i - 1, j + 1}) //up right
//		indexes = append(indexes, []int{i, j + 1})     // right
//		indexes = append(indexes, []int{i - 1, j})     // left
//		indexes = append(indexes, []int{i + 1, j})     //down
//		indexes = append(indexes, []int{i + 1, j - 1}) //down left
//		indexes = append(indexes, []int{i + 1, j + 1}) //down right
//
//		for _, v := range indexes {
//			if v[0] < 0 || v[0] < len(board)-1 || v[1] < 0 || v[1] < len(board[0])-1 {
//				continue
//			}
//			if board[v[0]][v[1]] == 0 {
//				O++
//			} else {
//				I++
//			}
//		}
//		org[0] = O
//		org[1] = I
//		return org
//	}
func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j, v := range board[i] {
			live := -v
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if x >= 0 && x < m && y >= 0 && y < n && board[x][y] > 0 {
						live++
					}
				}
			}
			if v == 1 && (live < 2 || live > 3) {
				board[i][j] = 2
			}
			if v == 0 && live == 3 {
				board[i][j] = -1
			}
		}
	}
	for i := 0; i < m; i++ {
		for j, v := range board[i] {
			if v == 2 {
				board[i][j] = 0
			}
			if v == -1 {
				board[i][j] = 1
			}
		}
	}
}
