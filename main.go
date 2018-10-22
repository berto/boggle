package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/berto/boggle/boggle"
	"github.com/berto/boggle/game"
)

const (
	startMessage   = "Find as many words as you can:"
	endMessage     = "Time's up!"
	gameTime       = 60 // in seconds
	dictionaryPath = "./dictionary/data/words.txt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	words := []string{}
	board := boggle.GenerateBoard()

	fmt.Println(game.PrintBoard(board))
	fmt.Println(startMessage)

	go func() {
		for scanner.Scan() {
			words = append(words, scanner.Text())
		}
	}()
	time.Sleep(gameTime * time.Second)

	fmt.Println(endMessage)
	fmt.Println(game.PrintScore(board, words))

	wordsMessage, err := game.PrintPossibleWords(board, dictionaryPath)
	if err != nil {
		fmt.Println("Error finding dictionary words")
	} else {
		fmt.Println(wordsMessage)
	}

	if scanner.Err() != nil {
		log.Fatal("Error scanning standard input")
	}
}
