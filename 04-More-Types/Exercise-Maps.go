package main
import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCounts(s string) map[string]int {

	wordMap := make(map[string]int)

	words := string.Fields(s)

	for _, word := range word {
		wordMap[word]++
	}

	return wordMap
}

func main() {
	wc.Test(WordCount)
}
