package print

import (
	"strconv"
	"strings"

	"github.com/berto/boggle/dictionary"
	"github.com/berto/boggle/game"
)

// Board displays the board string
func Board(board [][]byte) (boardString string) {
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

// Score displays the score result
func Score(board [][]byte, words []string, path string) string {
	newWords := []string{}
	for _, word := range words {
		finder := func(dictionaryWord string) bool {
			return dictionaryWord == strings.ToLower(word)
		}
		words, err := dictionary.FindDictionaryWords(finder, path)
		if err == nil && len(words) >= 1 {
			newWords = append(newWords, word)
		}
	}
	score := game.GradeBoard(board, newWords)
	return "Score: " + strconv.Itoa(score)
}

// PossibleWords displays possible dictionary words
func PossibleWords(board [][]byte, dictionaryPath string) (string, error) {
	finder := func(word string) bool {
		return game.FindWord(board, word)
	}
	words, err := dictionary.FindDictionaryWords(finder, dictionaryPath)
	if err != nil {
		return "", err
	}
	return "Possible words: " + strings.Join(words, " "), nil
}
