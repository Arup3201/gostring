package gostring

import (
	"strings"
	"unicode"
)

func CapitalizeWords(s string) string {
	var firstLetter string
	words := strings.Fields(s)
	cap := []string{}

	for _, w := range words {
		firstLetter = strings.ToUpper(string(w[0]))
		cap = append(cap, firstLetter+w[1:])
	}

	return strings.Join(cap, " ")
}

func CountVowelsConsonents(s string) (int, int) {
	const VOWELS = "aeiouAEIOU"

	cVowels, cConsonents := 0, 0

	for _, r := range s {
		if unicode.IsLetter(r) {
			if strings.ContainsRune(VOWELS, r) {
				cVowels++
			} else {
				cConsonents++
			}
		}
	}

	return cVowels, cConsonents
}

func IsPalindrome(s string) bool {
	return false
}

func ReverseString(s string) string {
	return ""
}
