package dictionary

import (
	"bufio"
	"os"
	"strings"

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
		word := strings.ToLower(scanner.Text())
		if finder(word) {
			words = append(words, word)
		}
	}
	return words, nil
}
