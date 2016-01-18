package main

import (
	"testing"
)

type Case struct {
	value string
	expected int
}

var cases = []Case {
	{"0", 0},
}

func TestValueOfIntegerToken(t *testing.T) {
	for _, c := range cases {
		tkn := token{INTEGER, c.value}
		if tkn.Int() != c.expected {
			t.Fail()
		}
	}
}
