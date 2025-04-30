package ast

type Node interface {
	node()
	String() string
}

type Stmt interface {
	Node
	stmt()
}

type Expr interface {
	Node
	expr()
}

type Lit interface {
	Expr
	lit()
}
