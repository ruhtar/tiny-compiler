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
	tokens := tokenizer(code) // Parsing

	fmt.Println(tokens)
}

func tokenizer(code string) []token {
	tokens := []token{}
	current := 0
	pendingParen := 0

	for current < len([]rune(code)) {
		char := string([]rune(code)[current])

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

		current++
	}

	if pendingParen%2 == 0 {
		panic("PARENTHESIS MISSING BRO")
	}

	return tokens
}
