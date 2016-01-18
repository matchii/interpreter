package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Notes here
// TODO read about types
// TODO read about interfaces
// TODO read about panic/recover
// TODO read about godoc

// Token types
const NONE = 0
const INTEGER = 1
const PLUS = 2
const MINUS = 3
const MUL = 4
const DIV = 5
const LPAREN = 6
const RPAREN = 7

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		text := strings.Split(strings.TrimRight(input, "\n"), "")
		if len(text) == 0 {
			continue
		}
		lxr := NewLexer(text)
		p := parser{nil, &lxr, GetTokenTypeNames()}
		fmt.Println(p.Expr())
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
	ttNames[LPAREN]  = "LPAREN"
	ttNames[RPAREN]  = "RPAREN"
	return ttNames
}

func IsDigit(char string) bool {
	isDigit, _ := regexp.MatchString("^[0-9]$", char)
	return isDigit
}
