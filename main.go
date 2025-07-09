package main

import "fmt"

type token struct {
	kind  string
	value string
}

// Parsing
// (add 2 (subtract 4 2))   =>   [{ type: 'paren', value: '(' }, ...]

func main() {
	code := "(add 2 (subtract 4 2))"
	// code := "( add )"
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
		}

		if char == ")" {
			tokens = append(tokens, token{
				kind:  "paren",
				value: char,
			})
			pendingParen--
		}

		isPossibleToHaveAnAdd := currentIndex+3 <= len(code) // if not checking this, go will panic in code[currentIndex:currentIndex+3] because it will be out of bounds
		if isPossibleToHaveAnAdd && code[currentIndex:currentIndex+3] == "add" {
			tokens = append(tokens, token{
				kind:  "name",
				value: "add",
			})
			currentIndex += 3
			continue
		}

		isPossibleToHaveASub := currentIndex+8 <= len(code) // if not checking this, go will panic in code[currentIndex:currentIndex+8] because it will be out of bounds
		if isPossibleToHaveASub && code[currentIndex:currentIndex+8] == "subtract" {
			tokens = append(tokens, token{
				kind:  "name",
				value: "subtract",
			})
			currentIndex += 8
			continue
		}

		currentIndex++
	}

	if pendingParen%2 != 0 {
		panic("PARENTHESIS MISSING BRO")
	}

	return tokens
}
