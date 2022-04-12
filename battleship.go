package main

func main() {
	board := createBoard(8, 8)

}

func startGame() [][]byte {
	board := createBoard(8, 8)

}

func createBoard(rows, cols int) [][]byte {

	board := make([][]byte, rows)
	for i := range board {
		board[i] = make([]byte, cols)
	}

	return board
}
