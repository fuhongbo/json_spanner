package parser

import (
	"errors"
	"strconv"
)

type lex struct {
	input     string
	scanner   *Scanner
	Query     Query
	ErrorInfo error
}

func NewLex(input string) *lex {
	return &lex{
		input:   input,
		scanner: newScanner(input),
	}
}

func (l *lex) Lex(lval *yySymType) int {
	typ, str := l.scanner.nextToken()

	switch typ {
	case 0:
		return 0
	case INTEGER:
		lval.int64, _ = strconv.ParseInt(str, 10, 64)
	case FLOAT:
		lval.float64, _ = strconv.ParseFloat(str, 64)
	//case STRING, NAME:
	//    lex.str = str
	case EQ, NEQ, LT, LTE, GT, GTE, AND, OR, ADD, DEC, TLIKE:
		lval.int = typ
	}
	lval.str = str

	return typ
}

func (l *lex) Error(err string) {
	l.ErrorInfo = errors.New(err)
}

func Parse(l *lex) Statements {
	yyParse(l)
	return l.Query.Statements
}
