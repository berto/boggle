package dictionary

import (
	"testing"

	"github.com/berto/boggle/boggle"
)

func TestFindPossibleWords(t *testing.T) {
	sampleBoard := [][]byte{
		{'a', 'b', 'c', 'd'},
		{'e', 't', 'e', 'f'},
		{'o', 's', 'a', 'r'},
		{'q', 'r', 'x', 'm'},
	}
	dictionaryPath := "./data/words_test.txt"
	fakeDictionaryPath := "./words_test.txt"
	finder := func(word string) bool {
		return boggle.FindWord(sampleBoard, word)
	}
	_, err := FindDictionaryWords(finder, fakeDictionaryPath)
	if err == nil {
		t.Errorf("Should throw error when file is not found: %s", err.Error())
	}

	words, err := FindDictionaryWords(finder, dictionaryPath)
	expected := []string{"bear", "beats", "team"}

	if err != nil || !testEqualWords(words, expected) {
		t.Errorf("Failed to find words from dictionary: expected %v to be %v", words, expected)
	}
}

func testEqualWords(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
