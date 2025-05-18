package leetcode

//@NOTE: Conditions
//@NOTE: 1 ALIVE
//@NOTE: nev keep 1 if 2 , 3 1's
//@NOTE: put 0 if 4 1's or  0 to 1 1's

//@NOTE: 0 DEAD
//@NOTE: put 1 if nev has 3 1's

func GameOfLife(board [][]int) {
	org := []int{}
	for i, _ := range board {
		for j, _ := range board[i] {
			org = heart(i, j, board)
			if board[i][j] == 0 {
				if org[0] > 3 {
					board[i][j] = 1
				}
			} else {
				if org[1] < 2 || org[1] > 3 {
					board[i][j] = 0
				}
			}
		}
	}
}

func heart(i, j int, board [][]int) (org []int) {
	return
}
