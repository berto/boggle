package boggle

import "testing"

func BenchmarkFindWord(b *testing.B) {
	testWords := []string{"testing", "is", "awesome"}
	for n := 0; n < b.N; n++ {
		board := generateBoard()
		gradeBoard(board, testWords)
	}
}
