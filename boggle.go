package boggle

import (
	"math/rand"
	"time"
)

const (
	alphabet  = "abcdefghijklmnopqrstuvwxyz"
	boardSize = 4
)

func generateLetter() byte {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(alphabet))
	return alphabet[randomIndex]
}

func generateBoard() [][]byte {
	board := make([][]byte, boardSize)
	for i, row := range board {
		row = make([]byte, boardSize)
		for j := range row {
			row[j] = generateLetter()
		}
		board[i] = row
	}
	return board
}
