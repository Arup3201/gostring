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
	lowStrWithoutSpace := strings.ReplaceAll(strings.ToLower(s), " ", "")
	strRunes := []rune(lowStrWithoutSpace)

	low, high := 0, len(strRunes)-1
	for low < high {
		if strRunes[low] != strRunes[high] {
			return false
		}

		low++
		high--
	}

	return true
}

func ReverseString(s string) string {
	low, high := 0, len(s)-1
	runes := []rune(s)

	for low < high {
		runes[low], runes[high] = runes[high], runes[low]
		low++
		high--
	}

	return string(runes)
}
