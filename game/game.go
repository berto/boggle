package game

import (
	"strconv"
	"strings"

	"github.com/berto/boggle/boggle"
	"github.com/berto/boggle/dictionary"
)

// PrintBoard displays the board string
func PrintBoard(board [][]byte) (boardString string) {
	for _, row := range board {
		for i, cell := range row {
			end := " "
			if i == len(row)-1 {
				end = "\n"
			}
			boardString += string(cell) + end
		}
	}
	return
}

// PrintScore displays the score result
func PrintScore(board [][]byte, words []string) string {
	score := boggle.GradeBoard(board, words)
	return "Score: " + strconv.Itoa(score)
}

// PrintPossibleWords displays possible dictionary words
func PrintPossibleWords(board [][]byte, dictionaryPath string) (string, error) {
	finder := func(word string) bool {
		return boggle.FindWord(board, word)
	}
	words, err := dictionary.FindDictionaryWords(finder, dictionaryPath)
	if err != nil {
		return "", err
	}
	return "Possible words: " + strings.Join(words, " "), nil
}
