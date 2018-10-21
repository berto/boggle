package boggle

import "testing"

const testString = "test"

func BenchmarkFindWord(b *testing.B) {
	board := generateBoard()
	for n := 0; n < b.N; n++ {
		findWord(board, testString)
	}
}
