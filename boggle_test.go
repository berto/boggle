package boggle

import (
	"reflect"
	"strings"
	"testing"
)

func TestBoggle(t *testing.T) {
	letter := generateLetter()

	assertLetter(t, letter)

	board := generateBoard()
	boardValue := reflect.ValueOf(board)
	if boardValue.Kind() != reflect.Slice {
		t.Errorf("Failed to generate board array")
	}
	if len(board) != boardSize && len(board[0]) != boardSize {
		t.Errorf("Failed to generate correct board size")
	}

	for _, row := range board {
		for _, letter = range row {
			assertLetter(t, letter)
		}
	}
}

func assertLetter(t *testing.T, letter byte) {
	if !strings.ContainsAny(alphabet, string(letter)) {
		t.Errorf("Letter not in the alphabet: %v", letter)
	}
}
