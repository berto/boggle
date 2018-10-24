package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/berto/boggle/game"
	"github.com/berto/boggle/print"
)

const (
	startMessage   = "Find as many words as you can:"
	endMessage     = "Time's up!"
	gameTime       = 60 // in seconds
	dictionaryPath = "../../dictionary/data/words.txt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	words := []string{}
	board := game.GenerateBoard()

	fmt.Println(print.Board(board))
	fmt.Println(startMessage)

	go func() {
		for scanner.Scan() {
			words = append(words, strings.ToLower(scanner.Text()))
		}
	}()
	time.Sleep(gameTime * time.Second)

	fmt.Println(endMessage)
	fmt.Println(print.Score(board, words, dictionaryPath))

	wordsMessage, err := print.PossibleWords(board, dictionaryPath)
	if err != nil {
		fmt.Println("Error finding dictionary words")
	} else {
		fmt.Println(wordsMessage)
	}

	if scanner.Err() != nil {
		log.Fatal("Error scanning standard input")
	}
}
