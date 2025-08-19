package main

import (
	"fmt"

	"github.com/Arup3201/gostring/gostring"
)

func main() {
	bookTitle := "the lord of the rings"

	fmt.Printf("Original: %s\n", bookTitle)
	fmt.Printf("Capitalize: %s\n", gostring.CapitalizeWords(bookTitle))

	vowels, consonents := gostring.CountVowelsConsonents(bookTitle)
	fmt.Printf("Vowels: %d, Consonents: %d\n", vowels, consonents)
	fmt.Printf("Is palindrome? %t\n", gostring.IsPalindrome(bookTitle))
	fmt.Printf("Reversed: %s\n", gostring.ReverseString(bookTitle))
}
