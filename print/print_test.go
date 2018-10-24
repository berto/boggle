package print

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

	got := Board(sampleBoard)

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
	words := []string{"lover", "apple", "team"}
	path := "../dictionary/data/words_test.txt"
	printedScore := Score(sampleBoard, words, path)
	expected := "Score: 2"
	if printedScore != expected {
		t.Errorf("Failed to print score: %s should equal %s", printedScore, expected)
	}
}

func TestPrintPossibleWords(t *testing.T) {
	sampleBoard := [][]byte{
		{'a', 'b', 'c', 'd'},
		{'e', 't', 'e', 'f'},
		{'o', 's', 'a', 'r'},
		{'q', 'r', 'x', 'm'},
	}
	expect := "Possible words: bear beats team"
	dictionaryPath := "../dictionary/data/words_test.txt"
	words, err := PossibleWords(sampleBoard, dictionaryPath)

	if err != nil || expect != words {
		t.Errorf("Failed to print words: expected %s to equal %s", words, expect)
	}
}
