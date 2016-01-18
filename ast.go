package main

import (
	//"fmt"
)

type node interface{
	//String() string
}

type binOp struct {
	left node
	op token
	right node
}

type num struct {
	t token
}

/*
func (node binOp) String() string {
	return fmt.Sprintf("OP(%s)", node.op.value)
}

func (node num) String() string {
	return fmt.Sprintf("int(value=%s)", node.t.value)
}
 */
