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

//////////////////////////////////////////////////////////////////////////////
//// Structures
//////////////////////////////////////////////////////////////////////////////

type token struct {
	tType int // because 'type' is reserved word in Go
	value string
}

type lexer struct {
	text []string
	pos int
	currentChar string
}

type interpreter struct {
	currentToken *token
	lexer *lexer
	// Names of token types, for descriptive error/debug messages
	ttNames map[int]string
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
		lexer := lexer{text, 0, text[0]}
		i := interpreter{nil, &lexer, GetTokenTypeNames()}
		fmt.Println(i.Expr())
	}
}

func GetTokenTypeNames() map[int]string {
	ttNames := make(map[int]string)
	ttNames[NONE]    = "NONE"
	ttNames[INTEGER] = "INTEGER"
	ttNames[PLUS]    = "PLUS"
	ttNames[MINUS]   = "MINUS"
	ttNames[MUL]     = "MUL"
	ttNames[DIV]     = "DIV"
	return ttNames
}

//////////////////////////////////////////////////////////////////////////////
//// Lexer
//////////////////////////////////////////////////////////////////////////////

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

//////////////////////////////////////////////////////////////////////////////
//// Interpreter
//////////////////////////////////////////////////////////////////////////////

// Expr evaluates expression
// Panics if try to divide by 0
// Returns result of calculation
func (i *interpreter) Expr() int {
	t := i.lexer.GetNextToken()
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
			"Token type mismatched, expected %s, got %s (value: %s)",
			i.ttNames[tokenType],
			i.ttNames[i.currentToken.tType],
			i.currentToken.value))
	}
	t := i.lexer.GetNextToken()
	i.currentToken = &t
}

//////////////////////////////////////////////////////////////////////////////
//// Token
//////////////////////////////////////////////////////////////////////////////

// Int converts token of type INTEGER to int value.
// Panics if called on non-INTEGER token.
func (t *token) Int() int {
	if t.tType != INTEGER {
		panic(fmt.Sprintf("This token is not INTEGER: %s", t.value))
	}
	r, _ := strconv.Atoi(t.value)
	return r
}

//////////////////////////////////////////////////////////////////////////////
//// Other functions
//////////////////////////////////////////////////////////////////////////////

func IsDigit(char string) bool {
	isDigit, _ := regexp.MatchString("^[0-9]$", char)
	return isDigit
}
