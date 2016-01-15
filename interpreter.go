package main

import (
	"fmt"
)

type interpreter struct {
	currentToken *token
	lexer *lexer
	// Names of token types, for descriptive error/debug messages
	ttNames map[int]string
}

// Expr evaluates expression
// expr : term ((PLUS | MINUS) term)*
// Panics if try to divide by 0
// Returns result of calculation
func (i *interpreter) Expr() int {
	if (i.currentToken == nil) {
    	t := i.lexer.GetNextToken()
    	i.currentToken = &t
	}

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

	return result
}

// Term
// term : factor ((MUL | DIV) factor)*
func (i *interpreter) Term() int {
	var result int
	result = i.Factor()

	for i.currentToken.tType == MUL || i.currentToken.tType == DIV {
		if i.currentToken.tType == MUL {
			i.Eat(MUL)
			result = result * i.Factor()
		} else if i.currentToken.tType == DIV {
			i.Eat(DIV)
			result = result / i.Factor()
		}
	}

	return result
}

// Factor
// factor : INTEGER | LPAREN expr RPAREN
func (i *interpreter) Factor() int {
	t := i.currentToken
	if t.tType == INTEGER {
		i.Eat(INTEGER)
		return t.Int()
	} else if t.tType == LPAREN {
		i.Eat(LPAREN)
		expr := i.Expr()
		i.Eat(RPAREN)
		return expr
	}
	// TODO use defer to make this non-terminating error
	panic(fmt.Sprintf("Found token %s where factor (INTEGER or LPAREN) expected", t.value))
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
