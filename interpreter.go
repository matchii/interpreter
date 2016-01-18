package main

type interpreter struct {}

func (i *interpreter) Visit(n node) int {
	switch n.(type) {
	case num:
		return i.VisitNum(n.(num))
	case binOp:
		return i.VisitOp(n.(binOp))
	default:
		panic("Visited unknown node type")
	}
}

func (i *interpreter) VisitNum(n num) int {
	return n.t.Int()
}

func (i *interpreter) VisitOp(n binOp) int {
	if n.op.tType == PLUS {
		return i.Visit(n.left) + i.Visit(n.right)
	} else if n.op.tType == MINUS {
		return i.Visit(n.left) - i.Visit(n.right)
	} else if n.op.tType == MUL {
		return i.Visit(n.left) * i.Visit(n.right)
	} else if n.op.tType == DIV {
		return i.Visit(n.left) / i.Visit(n.right)
	} else {
		panic("Visited unknown operator")
	}
}
