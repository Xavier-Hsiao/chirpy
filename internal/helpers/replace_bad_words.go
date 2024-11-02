package helpers

import "strings"

func ReplaceBadWords(text string, badWords map[string]struct{}) string {
	words := strings.Split(text, " ")
	for i, word := range words {
		lowerdWord := strings.ToLower(word)
		if _, exists := badWords[lowerdWord]; exists {
			words[i] = "****"
		}
	}
	cleanedWords := strings.Join(words, " ")
	return cleanedWords
}
