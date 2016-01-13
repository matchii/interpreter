package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	//"unicode/utf8"
)

// Token types
const NONE = 0
const INTEGER = 1
const PLUS = 2
const MINUS = 3
const MUL = 4
const DIV = 5

type token struct {
	tType int // because 'type' is reserved word in Go
	value string
}

type interpreter struct {
	text []string
	pos int
	currentToken *token
	currentChar string
}

// Notes here

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		text := strings.Split(strings.TrimRight(input, "\n"), "")
		if len(text) == 0 {
			continue
		}
		i := interpreter{text, NONE, nil, text[0]}
		fmt.Println(i.Expr())
	}
}

//// Interpreter methods

// Expr evaluates expression
// Panics if try to divide by 0
// Returns result of calculation
func (i *interpreter) Expr() int {
	t := i.GetNextToken()
	i.currentToken = &t

	var result int
	result = i.Term()

	for i.currentToken.tType == PLUS || i.currentToken.tType == MINUS {
		if i.currentToken.tType == PLUS {
			i.Eat(PLUS)
			result = result + i.Term()
		} else if i.currentToken.tType == MINUS {
			i.Eat(MINUS)
			result = result - i.Term()
		}
	}

	i.Eat(NONE)

	return result
}

// GetNextToken serves as a lexer
// Panics if finds unrecognized token
// Returns empty token when reaches end of text
func (i *interpreter) GetNextToken() token {
	//fmt.Printf("Current char: %s\n", i.currentChar)
	for i.currentChar != "" { // while there are still char in the input
		char := i.text[i.pos]
		if char == " " { // space, skip it
			i.Advance()
			continue
		}
		if IsDigit(char) {
			t := token{INTEGER, i.Integer()}
			return t
		}
		if char == "+" {
			t := token{PLUS, "+"}
			i.Advance()
			return t
		}
		if char == "-" {
			t := token{MINUS, "-"}
			i.Advance()
			return t
		}
		if char == "*" {
			t := token{MUL, "*"}
			i.Advance()
			return t
		}
		if char == "/" {
			t := token{DIV, "/"}
			i.Advance()
			return t
		}
		panic(fmt.Sprintf("Unknown token: %s at position %d", char, i.pos))
	}

	return token{NONE, ""}
}

func (i *interpreter) Term() int {
	t := i.currentToken
	i.Eat(INTEGER)
	return t.Int()
}

// Eat moves to new current token
// Panics if token type occurs to be different than expected
func (i *interpreter) Eat(tokenType int) {
	if i.currentToken.tType != tokenType {
		panic(fmt.Sprintf(
			"Token type mismatched, expected %d, got %d (value: %s)",
			tokenType,
			i.currentToken.tType,
			i.currentToken.value))
	}
	t := i.GetNextToken()
	i.currentToken = &t
}

// Advance moves pointer to next input position
func (i *interpreter) Advance() {
	i.pos++
	if i.pos > len(i.text)-1 {
		i.currentChar = ""
	} else {
		i.currentChar = i.text[i.pos]
	}
}

// Integer extracts integer value from input by concatenating subsequent digits
// Returns integer-like string
func (i *interpreter) Integer() string {
	var result []string
	for IsDigit(i.currentChar) {
		result = append(result, i.currentChar)
		i.Advance()
	}
	return strings.Join(result, "")
}

//// Token methods

// Int converts token of type INTEGER to int value.
// Panics if called on non-INTEGER token.
func (t *token) Int() int {
	if t.tType != INTEGER {
		panic(fmt.Sprintf("This token is not INTEGER: %s", t.value))
	}
	r, _ := strconv.Atoi(t.value)
	return r
}

func (t *token) Addition() bool {
	return t.tType == PLUS || t.tType == MINUS
}

//// Other functions

func IsDigit(char string) bool {
	isDigit, _ := regexp.MatchString("^[0-9]$", char)
	return isDigit
}
