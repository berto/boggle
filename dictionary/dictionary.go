package dictionary

import (
	"bufio"
	"os"

	"github.com/pkg/errors"
)

// FindDictionaryWords reads a dictionary file, runs a finder for each word, and returns matched words
func FindDictionaryWords(finder func(word string) bool, dictionaryPath string) ([]string, error) {
	file, err := os.Open(dictionaryPath)
	if err != nil {
		err = errors.Wrap(err, "unable to read dictionary")
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	words := []string{}
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) > 2 && finder(word) {
			words = append(words, word)
		}
	}
	return words, nil
}
