package boggle

import (
	"strings"
	"testing"
)

func TestBoggle(t *testing.T) {
	letter := generateLetter()

	if !strings.ContainsAny(alphabet, string(letter)) {
		t.Errorf("Failed to generate random letter: %v should be a letter in the alphabet", letter)
	}
}
