package boggle

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func generateLetter() byte {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(alphabet))
	return alphabet[randomIndex]
}
