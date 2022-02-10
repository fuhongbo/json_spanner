package parser

import (
	"log"
)

func (*SelectStatement) stmt() {}

type SelectStatement struct {
	Fields         Fields
	Source         *Source
	WhereCondition *ExprNode
}

func (stmt *SelectStatement) travel() {
	str := ""
	if len(stmt.Fields) != 0 {
		str = str + "SELECT " + stmt.Fields.str()
	}

	if stmt.WhereCondition != nil {
		str = str + " WHERE " + stmt.WhereCondition.str()
	}

	log.Println(str)
}
func (stmt *SelectStatement) info() SelectStatement {
	return *stmt
}
