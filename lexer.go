package main

import (
	"fmt"
	"strings"
)

type lexer struct {
	text []string
	pos int
	currentChar string
}

// GetNextToken serves as a lexer
// Panics if finds unrecognized token
// Returns empty token when reaches end of text
func (l *lexer) GetNextToken() token {
	//fmt.Printf("Current char: %s\n", i.currentChar)
	for l.currentChar != "" { // while there are still char in the input
		char := l.text[l.pos]
		if char == " " { // space, skip it
			l.Advance()
			continue
		}
		if IsDigit(char) {
			t := token{INTEGER, l.Integer()}
			return t
		}
		if char == "+" {
			t := token{PLUS, "+"}
			l.Advance()
			return t
		}
		if char == "-" {
			t := token{MINUS, "-"}
			l.Advance()
			return t
		}
		if char == "*" {
			t := token{MUL, "*"}
			l.Advance()
			return t
		}
		if char == "/" {
			t := token{DIV, "/"}
			l.Advance()
			return t
		}
		if char == "(" {
			t := token{LPAREN, "("}
			l.Advance()
			return t
		}
		if char == ")" {
			t := token{RPAREN, ")"}
			l.Advance()
			return t
		}
		panic(fmt.Sprintf("Unknown token: %s at position %d", char, l.pos))
	}

	return token{NONE, ""}
}

// Advance moves pointer to next input position
func (l *lexer) Advance() {
	l.pos++
	if l.pos > len(l.text)-1 {
		l.currentChar = ""
	} else {
		l.currentChar = l.text[l.pos]
	}
}

// Integer extracts integer value from input by concatenating subsequent digits
// Returns integer-like string
func (l *lexer) Integer() string {
	var result []string
	for IsDigit(l.currentChar) {
		result = append(result, l.currentChar)
		l.Advance()
	}
	return strings.Join(result, "")
}
