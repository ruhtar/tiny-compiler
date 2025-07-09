package main

import (
	"fmt"
	"unicode"
)

type token struct {
	kind  string
	value string
}

// Parsing
// (add 2 (subtract 4 2))   =>   [{ type: 'paren', value: '(' }, ...]

func main() {
	code := "(add 7 (subtract 32 21))"
	tokens := tokenizer(code) // Parsing

	for i, token := range tokens {
		fmt.Printf("[%d] %q\n", i, token)
	}

}

func tokenizer(code string) []token {
	tokens := []token{}
	currentIndex := 0
	pendingParen := 0

	for currentIndex < len([]rune(code)) {
		char := string([]rune(code)[currentIndex])

		if char == "(" {
			tokens = append(tokens, token{
				kind:  "paren",
				value: char,
			})
			pendingParen++
			currentIndex++
			continue
		}

		if char == ")" {
			tokens = append(tokens, token{
				kind:  "paren",
				value: char,
			})
			pendingParen--
			currentIndex++
			continue
		}

		if unicode.IsDigit(rune(char[0])) { //we want to return the numbers of multiple digits as one single token
			digits := getSequentialDigits(code[currentIndex:])
			tokens = append(tokens, token{
				kind:  "number",
				value: digits,
			})
			currentIndex += len(digits)
			continue
		}

		isPossibleToHaveAnAdd := currentIndex+3 <= len(code) // if not checking this, golang will panic in code[currentIndex:currentIndex+3] because it will be out of bounds
		if isPossibleToHaveAnAdd && code[currentIndex:currentIndex+3] == "add" {
			tokens = append(tokens, token{
				kind:  "name",
				value: "add",
			})
			currentIndex += 3
			continue
		}

		isPossibleToHaveASub := currentIndex+8 <= len(code) // if not checking this, golang will panic in code[currentIndex:currentIndex+8] because it will be out of bounds
		if isPossibleToHaveASub && code[currentIndex:currentIndex+8] == "subtract" {
			tokens = append(tokens, token{
				kind:  "name",
				value: "subtract",
			})
			currentIndex += 8
			continue
		}
		currentIndex++ // for whitespaces
	}

	if pendingParen%2 != 0 {
		panic("PARENTHESIS MISSING BRO")
	}
	return tokens
}

func getSequentialDigits(codeSubString string) string {
	digits := string(codeSubString[0])

	for i := 1; i < len(codeSubString); i++ {
		if unicode.IsDigit(rune(codeSubString[i])) {
			digits += string(codeSubString[i])
		} else {
			break
		}
	}

	return digits
}
