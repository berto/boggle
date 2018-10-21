package boggle

import (
	"reflect"
	"strings"
	"testing"
)

func TestBoggleGrader(t *testing.T) {
	words := []string{"pineapple", "lover", "secret", "love", "cab", "vampire", "ex", "fret"}
	scores := []int{11, 2, 3, 1, 1, 5, 0, 1}

	for i, word := range words {
		t.Run(word, func(t *testing.T) {
			score := scoreWord(word)
			if scores[i] != score {
				t.Errorf("Failed to score word: expected %v to equal %v", scores[i], score)
			}
		})
	}

	sampleBoard := [][]byte{
		{'a', 'b', 'c', 'd'},
		{'l', 'c', 't', 'f'},
		{'o', 'b', 'e', 'r'},
		{'q', 'v', 'x', 'm'},
	}

	got := gradeBoard(sampleBoard, words)
	want := 0 + 2 + 0 + 1 + 1 + 0 + 0 + 1

	if got != want {
		t.Errorf("Failed to score board: want %v to be %v", got, want)
	}
}

func TestBoggleGenerator(t *testing.T) {
	letter := generateLetter()

	assertLetter(t, letter)

	board := GenerateBoard()
	boardValue := reflect.ValueOf(board)
	if boardValue.Kind() != reflect.Slice {
		t.Error("Failed to generate board array")
	}
	if len(board) != boardSize && len(board[0]) != boardSize {
		t.Error("Failed to generate correct board size")
	}

	for _, row := range board {
		for _, letter = range row {
			assertLetter(t, letter)
		}
	}
}

func TestWordFinder(t *testing.T) {
	sampleBoardOne := [][]byte{
		{'a', 'b', 'c', 'd'},
		{'l', 'c', 't', 'f'},
		{'o', 'b', 'e', 'r'},
		{'q', 'v', 'x', 'm'},
	}
	wordOne := "lover"
	notWordOne := "hater"

	sampleBoardTwo := [][]byte{
		{'a', 'b', 'c', 'd'},
		{'l', 'h', 'g', 'f'},
		{'t', 'n', 'i', 'r'},
		{'q', 'e', 'r', 'm'},
	}
	wordTwo := "fighter"
	notWordTwo := "quitter"

	if FindWord(sampleBoardOne, wordOne) != true {
		t.Errorf("Failed to find word: %s", wordOne)
	}

	if FindWord(sampleBoardOne, notWordOne) == true {
		t.Errorf("Failed to not find word: %s", notWordOne)
	}

	if FindWord(sampleBoardTwo, wordTwo) != true {
		t.Errorf("Failed to find word: %s", wordTwo)
	}

	if FindWord(sampleBoardTwo, notWordTwo) == true {
		t.Errorf("Failed to not find word: %s", notWordTwo)
	}
}

func TestConverBoardAndSearch(t *testing.T) {
	sampleBoard := [][]byte{
		{'a', 'b', 'c', 'd'},
		{'a', 'b', 'c', 'd'},
		{'a', 'b', 'c', 'd'},
		{'a', 'b', 'c', 'd'},
	}
	result := convertBoard(sampleBoard)
	board := reflect.ValueOf(result)
	if board.Kind() != reflect.Slice && isNode(result[0]) {
		t.Error("Failed to convert board")
	}

	isThere := search(result, []byte("ab"))
	isNotThere := search(result, []byte("ax"))

	if isThere != true || isNotThere != false {
		t.Error("Failed to search word in board")
	}

	got := isNode("a")
	want := false

	got2 := isNode(node{})
	want2 := true

	if got != want || got2 != want2 {
		t.Error("Failed to check if something is node")
	}
}

func assertLetter(t *testing.T, letter byte) {
	if !strings.ContainsAny(alphabet, string(letter)) {
		t.Errorf("Letter not in the alphabet: %v", letter)
	}
}

func isNode(t interface{}) bool {
	switch t.(type) {
	case node:
		return true
	default:
		return false
	}
}
