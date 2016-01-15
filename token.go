package main

import (
	"fmt"
	"strconv"
)

type token struct {
	tType int // because 'type' is reserved word in Go
	value string
}

// Int converts token of type INTEGER to int value.
// Panics if called on non-INTEGER token.
func (t *token) Int() int {
	if t.tType != INTEGER {
		panic(fmt.Sprintf("This token is not INTEGER: %s", t.value))
	}
	r, _ := strconv.Atoi(t.value)
	return r
}
