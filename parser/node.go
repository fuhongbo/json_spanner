package parser

import (
	"fmt"
	"strconv"
)

type ExprNodeType int

const (
	Unknown     ExprNodeType = 0
	BinaryNode  ExprNodeType = 1
	FieldNode   ExprNodeType = 2
	IntegerNode ExprNodeType = 3
	FloatNode   ExprNodeType = 4
	StringNode  ExprNodeType = 5
)

type Statements []Statement
type Fields []*Field
type Source struct {
	Name string
}
type Field struct {
	Name        string
	Type        string
	ValueString string
	ValueInt    int64
	ValueFloat  float64
	Alias       string
}

type Statement interface {
	stmt()
	travel()
	info() SelectStatement
}

type Query struct {
	Statements Statements
}

type ExprNode struct {
	Type  ExprNodeType
	Left  *ExprNode
	Right *ExprNode

	FloVal float64
	IntVal int64
	StrVal string
	Name   string
	Op     int
}

func Travel(stmts Statements) {
	for _, stmt := range stmts {
		stmt.travel()
	}
}

func Info(stmt Statement) SelectStatement {
	return stmt.info()
}

func (fields Fields) str() string {
	str := ""
	for i, field := range fields {
		if i > 0 {
			str = str + ", "
		}

		str = str + field.str()
	}
	return str
}

func (field *Field) str() string {
	f := field.Name
	if field.Type == "INT" {
		f = fmt.Sprintf("%d", field.ValueInt)
	}
	if field.Type == "FLOAT" {

		f = fmt.Sprintf("%v", field.ValueFloat)
	}
	if field.Alias != "" {
		f = f + " AS " + field.Alias
	}
	return f
}

func (source *Source) str() string {
	return source.Name
}

func (node *ExprNode) str() string {
	switch node.Type {
	case FieldNode:
		return node.Name
	case StringNode:
		return node.StrVal
	case IntegerNode:
		return strconv.FormatInt(node.IntVal, 10)
	case FloatNode:
		return strconv.FormatFloat(node.FloVal, 'f', -1, 64)
	case BinaryNode:
		return fmt.Sprintf("(%s %s %s)", node.Left.str(), opToStr(node.Op), node.Right.str())
	}

	return "?"
}

func opToStr(op int) string {
	switch op {
	case EQ:
		return "="
	case NEQ:
		return "<>"
	case GT:
		return ">"
	case GTE:
		return ">="
	case LT:
		return "<"
	case LTE:
		return "<="
	case AND:
		return "AND"
	case OR:
		return "OR"
	case ADD:
		return "+"
	case DEC:
		return "-"
	case STAR:
		return "*"
	case COMMA:
		return ","
	case SEMICOLON:
		return ";"
	}
	return ""
}
