package main

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/berto/boggle/dictionary"
	"github.com/gin-gonic/gin"
)

const dictionaryPath = "./dictionary/data/words.txt"
const delay = 2 // in seconds

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r := gin.Default()
	r.GET("/dictionary", findRoute)
	r.GET("/delaytionary", func(c *gin.Context) {
		time.Sleep(delay * time.Second)
		findRoute(c)
	})
	r.Run(":" + port)
}

func findRoute(c *gin.Context) {
	word := c.Query("word")
	var v bool
	if word != "" {
		finder := func(dictionaryWord string) bool {
			return dictionaryWord == strings.ToLower(word)
		}
		words, err := dictionary.FindDictionaryWords(finder, dictionaryPath)
		if err == nil && len(words) >= 1 {
			v = true
		}
	}
	c.JSON(200, gin.H{
		"wordFound": strconv.FormatBool(v),
	})
}
