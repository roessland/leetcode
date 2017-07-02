package main

import (
	"regexp"
	"unicode"
)

func onlyHasLetters(letters map[rune]bool, word string) bool {
	for _, ch := range word {
		if !letters[unicode.ToLower(ch)] {
			return false
		}
	}
	return true
}

func runeSet(word string) map[rune]bool {
	set := map[rune]bool{}
	for _, ch := range word {
		set[ch] = true
	}
	return set
}

func findWords(words []string) []string {
	rows := []map[rune]bool{
		runeSet("qwertyuiop"),
		runeSet("asdfghjkl"),
		runeSet("zxcvbnm"),
	}
	oneRowWords := []string{}
	for _, word := range words {
		for _, row := range rows {
			if onlyHasLetters(row, word) {
				oneRowWords = append(oneRowWords, word)
				break
			}
		}
	}
	return oneRowWords
}

func findWordsRegex(words []string) []string {
	re := regexp.MustCompile("(?i)^([qwertyuiop]*|[asdfghjkl]*|[zxcvbnm]*)$")
	oneRowWords := []string{}
	for _, word := range words {
		if re.MatchString(word) {
			oneRowWords = append(oneRowWords, word)
		}
	}
	return oneRowWords
}

func main() {

}
