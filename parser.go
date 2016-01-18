package main

import (
	"fmt"
)

// FIXME I don't like how syntax errors are handled and displayed

type parser struct {
	currentToken *token
	lexer *lexer
	// Names of token types, for descriptive error/debug messages
	ttNames map[int]string
}

func (p *parser) Parse() node {
	return p.Expr()
}

// Expr evaluates expression
// expr : term ((PLUS | MINUS) term)*
// Panics if try to divide by 0
// Returns result of calculation
func (p *parser) Expr() node {
	defer func() { // TODO move to Parse()
		if r := recover(); r != nil {
			fmt.Println("Syntax error:", r)
		}
	}()
	if (p.currentToken == nil) {
    	t := p.lexer.GetNextToken()
    	p.currentToken = &t
	}

	left := p.Term()

	for p.currentToken.tType == PLUS || p.currentToken.tType == MINUS {
		prevToken := p.currentToken
		if p.currentToken.tType == PLUS {
			p.Eat(PLUS)
		} else if p.currentToken.tType == MINUS {
			p.Eat(MINUS)
		}

		right := p.Term()
		left = binOp{left, *prevToken, right}
	}

	return left
}

// Term
// term : factor ((MUL | DIV) factor)*
func (p *parser) Term() node {
	left := p.Factor()
	for p.currentToken.tType == MUL || p.currentToken.tType == DIV {
		prevToken := p.currentToken
		if p.currentToken.tType == MUL {
			p.Eat(MUL)
		} else if p.currentToken.tType == DIV {
			p.Eat(DIV)
		}
		right := p.Factor()
		left = binOp{left, *prevToken, right}
	}

	return left
}

// Factor
// factor : INTEGER | LPAREN expr RPAREN
func (p *parser) Factor() node {
	t := p.currentToken
	if t.tType == INTEGER {
		p.Eat(INTEGER)
		return num{*t}
	} else if t.tType == LPAREN {
		p.Eat(LPAREN)
		node := p.Expr()
		p.Eat(RPAREN)
		return node
	}
	panic(fmt.Sprintf(
		"Token '%s' (type %s) found at position %d when factor (INTEGER or LPAREN) expected",
		t.value,
		p.ttNames[t.tType],
		p.lexer.pos-1))
}

// Eat moves to new current token
// Panics if token type occurs to be different than expected
func (p *parser) Eat(tokenType int) {
	if p.currentToken.tType != tokenType {
		panic(fmt.Sprintf(
			"Token of type %s ('%s') found at position %d when %s expected",
			p.ttNames[p.currentToken.tType],
			p.currentToken.value,
			p.lexer.pos,
			p.ttNames[tokenType]))
	}
	t := p.lexer.GetNextToken()
	p.currentToken = &t
}
