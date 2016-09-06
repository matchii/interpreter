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
const (
	NONE = 0
	INTEGER = 1
	PLUS = 2
	MINUS = 3
	MUL = 4
	DIV = 5
	LPAREN = 6
	RPAREN = 7
)

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
		result, parseError := p.Parse()
		fmt.Println("result:", result, "parseError:", parseError)
		if parseError != nil {
			fmt.Println(parseError)
			continue
		}
		i := interpreter{}
		fmt.Println(i.Visit(result))
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
