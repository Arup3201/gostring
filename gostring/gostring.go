package gostring

import (
	"strings"
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
	return 0, 0
}

func IsPalindrome(s string) bool {
	return false
}

func ReverseString(s string) string {
	return ""
}
