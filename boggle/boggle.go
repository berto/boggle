package boggle

import (
	"math/rand"
	"time"
)

const (
	alphabet  = "abcdefghijklmnopqrstuvwxyz"
	boardSize = 4
)

type node struct {
	value     byte
	neighbors []*node
	checked   bool
}

// FindWord searches a word in a boggle board
func FindWord(board [][]byte, word string) bool {
	if len(word) < 3 {
		return false
	}
	wordMap := convertBoard(board)
	return search(wordMap, []byte(word))
}

// GenerateBoard creates a new boggle board
func GenerateBoard() [][]byte {
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

// GradeBoard finds matching words and returns a score
func GradeBoard(board [][]byte, words []string) int {
	score := 0
	for _, word := range words {
		if FindWord(board, word) {
			score += scoreWord(word)
		}
	}
	return score
}

func scoreWord(word string) int {
	if len(word) < 3 {
		return 0
	} else if len(word) < 5 {
		return 1
	} else if len(word) == 5 {
		return 2
	} else if len(word) == 6 {
		return 3
	} else if len(word) == 7 {
		return 5
	}
	return 11
}

func generateLetter() byte {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(alphabet))
	return alphabet[randomIndex]
}

func convertBoard(board [][]byte) []*node {
	newBoard := make([]*node, boardSize*boardSize)
	tempBoard := make([][]*node, boardSize)
	for i := range tempBoard {
		tempBoard[i] = make([]*node, boardSize)
	}
	for i, row := range board {
		for j, letter := range row {
			if tempBoard[i][j] == nil {
				tempBoard[i][j] = &node{value: letter}
			}
			newBoard[(i*boardSize)+j] = tempBoard[i][j]
			isNotTop := i-1 >= 0
			isNotFarLeft := j-1 >= 0
			isNotFarRight := j+1 < boardSize
			isNotBottom := i+1 < boardSize
			if isNotTop && isNotFarLeft {
				x, y := j-1, i-1
				addNeighbor(tempBoard[i][j], board[y][x], &tempBoard, x, y)
			}
			if isNotTop {
				x, y := j, i-1
				addNeighbor(tempBoard[i][j], board[y][x], &tempBoard, x, y)
			}
			if isNotTop && isNotFarRight {
				x, y := j+1, i-1
				addNeighbor(tempBoard[i][j], board[y][x], &tempBoard, x, y)
			}
			if isNotFarLeft {
				x, y := j-1, i
				addNeighbor(tempBoard[i][j], board[y][x], &tempBoard, x, y)
			}
			if isNotFarRight {
				x, y := j+1, i
				addNeighbor(tempBoard[i][j], board[y][x], &tempBoard, x, y)
			}
			if isNotFarLeft && isNotBottom {
				x, y := j-1, i+1
				addNeighbor(tempBoard[i][j], board[y][x], &tempBoard, x, y)
			}
			if isNotBottom {
				x, y := j, i+1
				addNeighbor(tempBoard[i][j], board[y][x], &tempBoard, x, y)
			}
			if isNotBottom && isNotFarRight {
				x, y := j+1, i+1
				addNeighbor(tempBoard[i][j], board[y][x], &tempBoard, x, y)
			}
		}
	}
	return newBoard
}

func addNeighbor(currentNode *node, letter byte, tempNode *[][]*node, x int, y int) {
	if (*tempNode)[y][x] == nil {
		(*tempNode)[y][x] = &node{value: letter}
	}
	if (*currentNode).neighbors == nil {
		(*currentNode).neighbors = []*node{}
	}
	(*currentNode).neighbors = append(currentNode.neighbors, (*tempNode)[y][x])
}

func search(nodes []*node, word []byte) bool {
	if len(word) == 0 {
		return true
	}
	for i := range nodes {
		if nodes[i].value == word[0] && !nodes[i].checked {
			nodes[i].checked = true
			if search(nodes[i].neighbors, word[1:]) {
				return true
			}
		}
	}
	return false
}
