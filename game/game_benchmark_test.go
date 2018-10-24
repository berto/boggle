package game

import "testing"

func BenchmarkGradeBoard(b *testing.B) {
	testWords := []string{"testing", "is", "awesome"}
	for n := 0; n < b.N; n++ {
		board := GenerateBoard()
		GradeBoard(board, testWords)
	}
}
