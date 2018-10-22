package game

import (
	"testing"
)

func TestPrintGame(t *testing.T) {
	sampleBoard := [][]byte{
		{'a', 'b', 'c', 'd'},
		{'l', 'c', 't', 'f'},
		{'o', 'b', 'e', 'r'},
		{'q', 'v', 'x', 'm'},
	}
	want := `a b c d
l c t f
o b e r
q v x m
`

	got := PrintBoard(sampleBoard)

	if got != want {
		t.Errorf("Failed to print board: %s should equal %s", got, want)
	}
}

func TestPrintScore(t *testing.T) {
	sampleBoard := [][]byte{
		{'a', 'b', 'c', 'd'},
		{'l', 'c', 't', 'f'},
		{'o', 'b', 'e', 'r'},
		{'q', 'v', 'x', 'm'},
	}
	words := []string{"lover", "apple"}
	if PrintScore(sampleBoard, words) != "Score: 2" {
		t.Error("Failed to print score")
	}
}

func TestPrintPossibleWords(t *testing.T) {
	sampleBoard := [][]byte{
		{'a', 'b', 'c', 'd'},
		{'e', 't', 'e', 'f'},
		{'o', 's', 'a', 'r'},
		{'q', 'r', 'x', 'm'},
	}
	expect := "Possible words: bear beats"
	dictionaryPath := "../dictionary/data/words_test.txt"
	words, err := PrintPossibleWords(sampleBoard, dictionaryPath)

	if err != nil || expect != words {
		t.Errorf("Failed to print words: expected %s to equal %s", words, expect)
	}
}
