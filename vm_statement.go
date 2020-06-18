package jack_compiler

type statementCategory int
const (
	ifSc statementCategory = iota
	whileSc
	retSc
	doSc
	letSc
)

type Statement interface {
	category() statementCategory
}

type termCategory int

const(
	constantTerm termCategory = iota
	referenceTerm
	unaryTerm
	expressionTerm
	subCallTerm
)

type Term interface {
	category() termCategory
}

type ConstTerm struct {
	ttype TokenType //keyword, string, integer
	val interface{}
}

func (ct ConstTerm) category() termCategory {
	return constantTerm
}

type UnaryTerm struct {
	operator string
	term 	 Term
}

func (ut UnaryTerm) category() termCategory {
	return unaryTerm
}

type ReferenceTerm struct {
	varName string
	index expression
}

func (rt ReferenceTerm) category() termCategory {
	return referenceTerm
}

type expression struct {
	terms []Term
	operations []string
}

func (exp expression) category() termCategory {
	return expressionTerm
}

type ifStatement struct {
	condition expression
	statements []Statement
	elseStatements []Statement
}

func (is ifStatement) category() statementCategory {
	return ifSc
}

type doStatement struct {
	action subroutineCall
}

func (ds doStatement) category() statementCategory {
	return doSc
}

type letStatement struct {
	target ReferenceTerm
	expression expression
}

func (ls letStatement) category() statementCategory {
	return letSc
}

type retStatement struct {
	expression expression
}

func (rs retStatement) category() statementCategory {
	return retSc
}


